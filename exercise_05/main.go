package main

import (
	"fmt"
	"net/http"
)

func agentHandler(w http.ResponseWriter, r *http.Request) {

	agent := r.Header.Get("User-Agent")

	if agent == "" {
		agent = "Unknown Client"

	}
	fmt.Fprintf(w, "You are visiting us using: %s", agent)
}

func main() {
	http.HandleFunc("/agent", agentHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
