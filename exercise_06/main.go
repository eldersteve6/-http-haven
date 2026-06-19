package main

import (
	"fmt"
	"net/http"
)

func legacyHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(
		w,
		r,
		"/v2",
		http.StatusMovedPermanently,
	)
}

func v2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2")
}

func main() {
	http.HandleFunc("/legacy", legacyHandler)

	http.HandleFunc("/v2", v2Handler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
