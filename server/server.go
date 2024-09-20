package server

import (
	"fmt"
	"net/http"
)

// StartServer starts the HTTP server
func StartServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}