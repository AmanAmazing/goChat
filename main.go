package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AmanAmazing/goChat/routes"
	"github.com/AmanAmazing/goChat/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	utils.SetTokenAuth(os.Getenv("JWT_SECRET_KEY"))
	err := utils.InitDB()
	if err != nil {
		log.Fatalf("Database initilisation error: %v", err)
	}
	testData := false
	if testData {
		utils.TestDB()
	}

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// creating a fileserver
	dir := http.Dir("./assets")
	fs := http.FileServer(dir)
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	r.Mount("/", routes.UserRouter())

	r.Get("/wscanvas", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		err = conn.WriteMessage(websocket.TextMessage, []byte("This is binary encoded message"))
		if err != nil {
			log.Println(err)
			return
		}

	})

	http.ListenAndServe(os.Getenv("PORT"), r)
}
