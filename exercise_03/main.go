package main

import (
	"fmt"
	"io"
	"net/http"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request with text to count words")
		return
	}

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		text := string(body)

		fmt.Fprint(w, len(text))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/count", countHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}
