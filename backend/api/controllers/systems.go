package controllers

import (
	"encoding/json"
	"net/http"
)

// GetVersion - Get API version
func GetVersion(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"version": "1.0"})
}

// Ping - health check
func Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
