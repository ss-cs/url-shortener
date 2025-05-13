// main.go
package main

import (
	"log"
	"net/http"
	"url-shortener/auth"
	"url-shortener/handler"
	"url-shortener/store"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	store.InitMongo()

	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("/register", handler.RegisterPage)
	http.HandleFunc("/login", handler.LoginPage)
	http.HandleFunc("/shorten", handler.Secure(handler.ShortenURL))

	http.HandleFunc("/api/register", handler.RegisterUser)
	http.HandleFunc("/api/login", handler.LoginUser)

	http.HandleFunc("/static/", handler.StaticFiles)
	http.HandleFunc("/", handler.ResolveURL)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
