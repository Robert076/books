package main

import (
	"encoding/json"
	"log"
	"net/http"

	model "github.com/Robert076/books.git/src/model"
)

func main() {
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(model.Books); err != nil {
			http.Error(w, "Error encoding books", http.StatusInternalServerError)
			return
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server.")
	}
}
