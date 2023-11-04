package main

import (
	"fmt"
	"log"
	"os"

	// "mime/multipart"
	"net/http"
	// "path/filepath"
)

func main() {
	http.HandleFunc("/samples", streamHandler)

	err := http.ListenAndServe("127.0.0.1:8080", nil)

	if err != nil {
		log.Fatalf("Failed with error: %s", err)
	}
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !r.URL.Query().Has("v") {
		http.Error(w, "Parameter \"v\" required", http.StatusBadRequest)
		return
	}

	fileName := r.URL.Query().Get("v")

	data, err := os.ReadFile(fmt.Sprintf("./samples/%s", fileName))

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read file with error: %s", err), http.StatusBadRequest)
		return
	}

	w.Write(data)
}
