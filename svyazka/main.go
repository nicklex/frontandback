package main

import (
	"fmt"
	"net/http"
)

func main() {
	// API
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"message": "привет с бекенда ()"}`)
	})

	// для css чтобы стили были
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server is running on port 8000...")
	http.ListenAndServe(":8000", nil)
}
