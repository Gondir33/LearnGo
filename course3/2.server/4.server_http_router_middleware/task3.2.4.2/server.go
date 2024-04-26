// Пример кода для создания сервера на go-chi

package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func main() {
	r := chi.NewRouter()

	// Применение middleware для логирования с помощью zap logger
	r.Use(LoggerMiddleware)

	// Здесь можно добавить ваши маршруты с различными методами
	r.HandleFunc("/1", handleRoute1)

	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, _ := zap.NewProduction()
		logger.Info("The info about clinet",
			zap.String("URL", r.URL.String()),
			zap.String("METHOD", r.Method),
			zap.String("IP ADDRESS", r.RemoteAddr))
	})
}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	// Обработка маршрута 1
}
