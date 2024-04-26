// Пример кода для создания сервера на go-chi
package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.HandleFunc("/1", handleRoute1)
	r.HandleFunc("/2", handleRoute2)
	r.HandleFunc("/3", handleRoute3)

	http.ListenAndServe(":8080", r)

}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	// Обработка маршрута 1
}

func handleRoute2(w http.ResponseWriter, r *http.Request) {
	// Обработка маршрута 2
}

func handleRoute3(w http.ResponseWriter, r *http.Request) {
	// Обработка маршрута 3
}
