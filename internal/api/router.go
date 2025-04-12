package api

import (
	"net/http"

	"go-rest-api/internal/handlers"
)

func NewRouter(itemHandler *handlers.ItemHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			itemHandler.GetAllItems(w, r)
		case http.MethodPost:
			itemHandler.CreateItem(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			itemHandler.GetItemById(w, r)
		case http.MethodDelete:
			itemHandler.DeleteItem(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
