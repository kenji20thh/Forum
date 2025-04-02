package main

import (
	"fmt"
	"forum/server"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting forum server on: 8080...")
	log.Fatal(http.ListenAndServe("8089", server.NewRouter()))
}
