package main

import (
	"fmt"
	"log"
	"net/http"
)

type handler struct{}

func (m handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Golang")
}

func main() {

	h := handler{}

	err := http.ListenAndServe(":8000", h)
	log.Fatal(err)
}
