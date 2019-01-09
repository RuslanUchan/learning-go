package main

import (
	"fmt"
	"log"
	"net/http"
)

// req 1 => req 2

func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s request %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	h := http.NewServeMux() // http.Handler interfaces

	h.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About Page")
	})

	h.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Contact Page")
	})

	h.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Blog Page")
	})

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Index Page")
	})

	// logger middleware
	lr := logRequest(h)

	log.Println("Server running on port 8000")

	err := http.ListenAndServe(":8000", lr)
	if err != nil {
		log.Fatal(err)
	}
}
