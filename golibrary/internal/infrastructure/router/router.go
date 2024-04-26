package router

import (
	"golibrary/internal/infrastructure/component"
	"golibrary/internal/modules"
	"golibrary/internal/router"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(controllers *modules.Controllers, components *component.Components) *chi.Mux {
	r := chi.NewRouter()
	r.Use(SwaggerHandler)
	r.Use(middleware.Logger)
	r.Mount("/", router.NewApiRouter(controllers, components))
	return r
}

func SwaggerHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/docs") {
			http.ServeFile(w, r, "/docs/swagger.json")
		} else if strings.HasPrefix(r.URL.String(), "/swagger") {
			swaggerUI(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
