package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.HandleFunc("/group{groupid}/{id}", groupHandler)
	http.ListenAndServe(":80", r)
}

func groupHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	groupid := chi.URLParam(r, "groupid")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Group "+groupid+" Привет, мир "+id)
}

/*
func main() {
	r := chi.NewRouter()

	r.Route("/group1", func(r chi.Router) {
		r.HandleFunc("/1", group1Handler1)
		r.HandleFunc("/2", group1Handler1)
		r.HandleFunc("/3", group1Handler1)
	})

	r.Route("/group2", func(r chi.Router) {
		r.HandleFunc("/1", group1Handler1)
		r.HandleFunc("/2", group1Handler1)
		r.HandleFunc("/3", group1Handler1)
	})

	r.Route("/group3", func(r chi.Router) {
		r.HandleFunc("/1", group1Handler1)
		r.HandleFunc("/2", group1Handler1)
		r.HandleFunc("/3", group1Handler1)
	})

}

func group1Handler1(w http.ResponseWriter, r *http.Request) {

}
*/
