package server

import (
	"fmt"
	"log"
	"net/http"
)

// Up .
func Up() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Github actions CI/CD ")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to test")
	})

	log.Fatal(http.ListenAndServe(":1323", nil))
}
