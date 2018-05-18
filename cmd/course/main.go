package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", servePresentation)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func servePresentation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Dear!")
}