package router

import (
	"url-shortener/handlers"
	"url-shortener/storage"

	"github.com/gorilla/mux"
)

func NewRouter(store *storage.MongoStorage) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", handlers.CreateShortURLHandler(store)).Methods("POST")
	r.HandleFunc("/{short}", handlers.RedirectURLHandler(store)).Methods("GET")
	return r
}
