package main

import (
	"fmt"
	"log"
	"net/http"

	"./handler"
	"./model"
)

func main() {
	pool := model.InitPool()
	conn := pool.Get()

	defer conn.Close()

	err := model.Ping(conn)
	if err != nil {
		fmt.Println(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.UrlHandler)
	log.Println("Localhost is running on port 8000")
	http.ListenAndServe(":8000", mux)
}
