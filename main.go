package main

import (
	"log"
	"net/http"
	"url-shortener/handler"
)

func main(){
	http.HandleFunc("/shorten", handler.ShortenURLHandler)
	http.HandleFunc("/", handler.RedirectHandler)

	log.Println("server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}