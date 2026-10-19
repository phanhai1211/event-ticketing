package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	mux.HandleFunc("GET /version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := VersionResponse{Version: "0.1.0"}
		json.NewEncoder(w).Encode(response)
	})

	log.Println("Server chạy tại http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
