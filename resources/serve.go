package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	if err := http.ListenAndServe(":9929", nil); err != nil {
		log.Fatal(err)
	}
}
