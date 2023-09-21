package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.FileServer(http.Dir("./docs"))
	log.Fatal(http.ListenAndServe(":7353", server))
}
