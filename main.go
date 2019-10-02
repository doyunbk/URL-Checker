package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/url-checker/handler"
	redis "github.com/url-checker/model"
)

func main() {
	pool := redis.InitPool()
	conn := pool.Get()

	defer conn.Close()

	err := redis.Ping(conn)
	if err != nil {
		fmt.Println(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.UrlHandler)
	log.Println("Localhost is running on port 8000")
	http.ListenAndServe("localhost:8000", mux)
}
