package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hey there , engineering flix")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hellow from redbus")
	})

	http.HandleFunc("/hii", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(r.URL.Path)
		fmt.Fprintf(w, "Hello,%q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":6060", nil))
}
