package main

import (
	"fmt"
	"log"
	"net/http"
)

// UrlHandler process HTTP request
func UrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url Handler")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UrlHandler)
	log.Println("Localhost is running on port 8000")
	http.ListenAndServe("localhost:8000", mux)
}
