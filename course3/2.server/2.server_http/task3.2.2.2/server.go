package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	port, exists := os.LookupEnv("PORT")

	if exists {
		http.HandleFunc("/", Handler)
		log.Printf("Сервер запущен на порту %s", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello world")
}
