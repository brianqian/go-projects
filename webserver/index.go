package main

import (
	"fmt"
	"log"
	"net/http"
)

var PORT = "8080"

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health check!")
	path := r.URL.Path
	if path != "/health" {
		msg := fmt.Sprintf("404 %s not found", path)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Form hit!")
	path := r.URL.Path
	if path != "/form" {
		msg := fmt.Sprintf("404 %s not found", path)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/health", healthCheck)

	fmt.Printf("Starting server at port %s", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}
