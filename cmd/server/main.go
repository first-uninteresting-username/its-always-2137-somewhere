package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloYouRequest struct {
	You string `json:"you"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello-world", func(w http.ResponseWriter, r *http.Request) {
		writeResponse(w, http.StatusOK, "hello, world")
	})

	mux.HandleFunc("POST /hello-you", func(w http.ResponseWriter, r *http.Request) {
		var req helloYouRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		writeResponse(w, http.StatusOK, fmt.Sprintf("hello, %s", req.You))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func writeResponse(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writing response: %v", err)
	}
}
