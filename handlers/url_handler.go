package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"url-shortener/models"
	"url-shortener/storage"

	"github.com/gorilla/mux"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateShortURLHandler(store *storage.MongoStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Original string `json:"original"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		short := randSeq(6)
		url := models.URL{
			Original: req.Original,
			Short:    short,
		}
		if err := store.Save(url); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(url)
	}
}

func RedirectURLHandler(store *storage.MongoStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		short := vars["short"]

		url, exists := store.Get(short)
		if !exists {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, url.Original, http.StatusMovedPermanently)
	}
}
