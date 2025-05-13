package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"url-shortener/auth"
	"url-shortener/models"
	"url-shortener/store"
)

func Secure(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		_, err = auth.ValidateJWT(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		h(w, r)
	}
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, _ := template.ParseFiles("static/shorten.html")
		tmpl.Execute(w, nil)
		return
	}

	r.ParseForm()
	url := r.FormValue("url")
	code, _ := store.SaveURL(url)
	fmt.Fprintf(w, "Shortened: <a href=\"/%s\">/%s</a>", code, code)
}
