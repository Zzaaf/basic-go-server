package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	println("Server listen on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

	http.HandleFunc("/", mainPage)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	welcomeText := "Hello, Go!"
	w.Write([]byte(welcomeText))
}
