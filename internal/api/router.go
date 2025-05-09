package api

import (
	"net/http"

	"go-rest-api/internal/dependencies"
)

func NewRouter(deps *dependencies.Dependencies) http.Handler {
	mux := http.NewServeMux()

	// Register routes by domain
	registerItemRoutes(mux, deps)
	// registerUserRoutes(mux, deps)

	return mux
}

func registerItemRoutes(mux *http.ServeMux, deps *dependencies.Dependencies) {
	mux.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if id := r.URL.Query().Get("id"); id != "" { // If an 'id' is provided as a query parameter
				deps.ItemHandler.GetItemById(w, r) // Fetch item by ID
			} else {
				deps.ItemHandler.GetAllItems(w, r) // Fetch all items if no 'id' is provided
			}
		case http.MethodPost:
			deps.ItemHandler.CreateItem(w, r)
		case http.MethodDelete:
			deps.ItemHandler.DeleteItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

// func registerUserRoutes(mux *http.ServeMux, deps *dependencies.Dependencies) {
// 	mux.Handle("/users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			deps.UserHandler.GetAll(w, r)
// 		case http.MethodPost:
// 			deps.UserHandler.Create(w, r)
// 		default:
// 			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		}
// 	}))
// }
