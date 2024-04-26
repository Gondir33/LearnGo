package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

const port = "80"

func main() {
	r := chi.NewRouter()

	r.HandleFunc("/1", firstHandler)
	r.HandleFunc("/2", secondHandler)
	r.HandleFunc("/3", thirdHandler)

	r.HandleFunc("/", defaultHandler)

	http.ListenAndServe(":"+port, r)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not Found")
}

func firstHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello world")
}

func secondHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello world 2")
}

func thirdHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello world 3")
}
