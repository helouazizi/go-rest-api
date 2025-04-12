package utils

import (
	"encoding/json"
	"net/http"
)

// Helper function to standardize error responses
func RespondWithError(w http.ResponseWriter, statusCode int, message, details string) {
	response := map[string]interface{}{
		"status":  "error",
		"message": message,
		"details": details,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Helper function to standardize JSON responses
func RespondWithJSON(w http.ResponseWriter, statusCode int, result any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(result)
}
