package handler

import (
	"encoding/json"
	"net/http"
	"url-shortener/auth"
	"url-shortener/store"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	err := store.CreateUser(email, password)
	if err != nil {
		http.Error(w, "User exists", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	if err := store.ValidateUser(email, password); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, _ := auth.GenerateJWT(email)
	http.SetCookie(w, &http.Cookie{Name: "token", Value: token, HttpOnly: true})
	http.Redirect(w, r, "/shorten", http.StatusSeeOther)
}
