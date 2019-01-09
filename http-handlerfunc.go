package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Golang")
	})

	log.Println("Server running on port 8000")

	err := http.ListenAndServe(":8000", h)
	if err != nil {
		log.Fatal(err)
	}
}
