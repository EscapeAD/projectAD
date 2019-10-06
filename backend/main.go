package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/version", getVersion)
	fmt.Println("Server online")
	http.ListenAndServe(":3000", r)
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"version": "1.0"})
}
