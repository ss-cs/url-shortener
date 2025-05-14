package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"url-shortener/model"
	"url-shortener/storage"
	"url-shortener/utils"
)

var store = storage.NewMemoryStore()

func ShortenURLHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post allowed", http.StatusMethodNotAllowed)
		return
	}
	var req model.URLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == ""{
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	code := utils.GenerateShortCode()
	store.Save(code, req.URL)

	resp := model.URLResponse{ShortURL : "http://localhost:8080/"+code}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(resp)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request){
	code := strings.TrimPrefix(r.URL.Path, "/")
	originalURL, found := store.Get(code)

	if !found {
		http.NotFound(w,r)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}