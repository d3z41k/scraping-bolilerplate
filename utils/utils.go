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

// Result returns a map result
func Result(item string, status bool, message string) map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		item: map[string]interface{}{
			"status":  status,
			"message": message,
		},
	}
}

// RsultMessage returns a result message
func RsultMessage(item string, status bool) string {
	m := "Ok"

	switch item {
	case "name":
		if status != true {
			m = "Metatag not found"
		}
	case "content":
		if status != true {
			m = "Invalid metatag content"
		}
	case "phrase":
		if status != true {
			m = "Search phrase not found"
		}
	}

	return m
}

// Respond add JSON header and encode data
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
