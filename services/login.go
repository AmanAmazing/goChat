package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AmanAmazing/goChat/models"
	"github.com/AmanAmazing/goChat/utils"
	"github.com/jackc/pgx/v5"
)

func PostLogin(username, password string) (string, error) {
	var user models.PostLogin
	err := utils.Db.QueryRow(context.Background(), "SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrInvalidCredentials
		}
		return "", ErrDatabaseError
	}
	fmt.Println(user)
	match := utils.CheckPasswordMatch(user.Password, password)
	if !match {
		return "", ErrInvalidCredentials
	}
	// Generate a JWT token
	claims := map[string]interface{}{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Sign the token with a secret key
	tokenString, err := utils.GenerateJWT(claims)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}
	return tokenString, nil

}
