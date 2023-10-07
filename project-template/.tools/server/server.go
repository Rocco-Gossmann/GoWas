package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.FileServer(http.Dir("./docs"))
	log.Println("starting Server http://localhost:7353")
	log.Fatal(http.ListenAndServe(":7353", server))
}
