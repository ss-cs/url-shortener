package handler

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/register", http.StatusSeeOther)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/register.html")
	tmpl.Execute(w, nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/login.html")
	tmpl.Execute(w, nil)
}

func StaticFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
