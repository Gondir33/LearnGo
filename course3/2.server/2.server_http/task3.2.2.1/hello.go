package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello world")
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(portNumber, nil)
}
