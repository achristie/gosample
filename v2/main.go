package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPersonServer(NewStore())
	log.Fatal(http.ListenAndServe(":8080", server))
}
