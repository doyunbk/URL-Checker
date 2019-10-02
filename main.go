package main

import (
	"fmt"
	"log"
	"net/http"

	redis "github.com/url-checker/model"
)

// UrlHandler process HTTP request
func UrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url Handler")
}

func main() {
	pool := redis.InitPool()
	conn := pool.Get()

	defer conn.Close()

	err := redis.Ping(conn)
	if err != nil {
		fmt.Println(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", UrlHandler)
	log.Println("Localhost is running on port 8000")
	http.ListenAndServe("localhost:8000", mux)
}
