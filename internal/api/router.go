package api

import (
	"net/http"

	"go-rest-api/internal/dependencies"
)

func NewRouter(deps *dependencies.Dependencies) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			deps.ItemHandler.GetAllItems(w, r)
		case http.MethodPost:
			deps.ItemHandler.CreateItem(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			deps.ItemHandler.GetItemById(w, r)
		case http.MethodDelete:
			deps.ItemHandler.DeleteItem(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
