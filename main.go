package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/url-checker/handler"
	"github.com/url-checker/model"
)

func main() {
	pool := model.InitPool()
	model.SeedData()
	conn := pool.Get()

	defer conn.Close()

	err := model.Ping(conn)
	if err != nil {
		fmt.Println(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.UrlHandler)
	log.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", mux)
}
