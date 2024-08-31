package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is missing")
	}

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Printf("JSON Decode error: %s\n", err)
	}

	return err
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("JSON Encode error: %s\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func ErrorJSON(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
