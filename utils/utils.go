package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var Db *pgxpool.Pool
var TokenAuth *jwtauth.JWTAuth

func HashPassword(plainPass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPass), 14)
	return string(hashedPassword), err
}

func CheckPasswordMatch(hashedPassword, enteredPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
	return err == nil
}
func SetTokenAuth(secretKey string) error {
	if secretKey == "" {
		return errors.New("JWT_SECRET_KEY environment variable not found")
	}
	TokenAuth = jwtauth.New("HS256", []byte(secretKey), nil)
	return nil
}

func GenerateJWT(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func InitDB() error {
	var err error
	dsn := fmt.Sprintf("%s://%s:%s@localhost:%s/%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	Db, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return err
	}
	return Db.Ping(context.Background())
}

func TestDB() {
	test_data := "test_data.sql"
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("%s://%s:%s@localhost:%s/%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	script, err := os.ReadFile(test_data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read SQL script: %v\n", err)
		os.Exit(1)
	}
	_, err = conn.Exec(context.Background(), string(script))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing SQL script: %v\n", err)
		os.Exit(1)
	}

	log.Println("test data inserted successfully")
}
