package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// each request called a handler
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}
