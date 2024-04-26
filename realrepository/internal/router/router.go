package router

import (
	"net/http"
	"repository/internal/controller"
	"repository/internal/infrastructure/responder"
	swag "repository/internal/infrastructure/router"
	"repository/internal/storage"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ptflp/godecoder"
	"go.uber.org/zap"
)

func NewApiHandler(Storage storage.UserRepository) http.Handler {
	r := chi.NewRouter()

	logger, _ := zap.NewProduction()

	userer := controller.NewUserer(Storage, responder.NewResponder(godecoder.NewDecoder(), logger))

	r.Use(middleware.Logger)
	r.Use(SwaggerHandler)

	r.Group(func(r chi.Router) {
		r.Post("/api/users/create", userer.CreateUserhandler)
		r.Post("/api/users/update", userer.UpdateUserhandler)
		r.Post("/api/users/delete", userer.DeleteUserhandler)
		r.Post("/api/users/list", userer.ListUserhandler)

		r.Get("/api/users/{id}", userer.GetByIdUserhandler)
	})

	return r
}

func SwaggerHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/docs") {
			http.ServeFile(w, r, "/docs/swagger.json")
		} else if strings.HasPrefix(r.URL.String(), "/swagger") {
			swag.SwaggerUI(w, r)
		} else if strings.HasPrefix(r.URL.String(), "/api") {
			next.ServeHTTP(w, r)
		}
	})
}
