package main

import "net/http"

func main() {
	// Simple Http Server
	http.ListenAndServe(":8080", nil)
}
