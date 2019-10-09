package utils

import (
	"encoding/json"
	"net/http"
)

// Message returns a map message
func Message(code int, status bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"status":  status,
		"message": message,
	}
}

// Respond add JSON header and encode data
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
