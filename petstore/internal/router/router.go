package router

import (
	"net/http"
	swag "petstore/internal/infrastructure/router"
	"petstore/internal/modules"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
)

func NewApiHandler(handlers *modules.Handlers, token *jwtauth.JWTAuth) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(SwaggerHandler)

	r.Group(func(r chi.Router) {
		r.Post("/user", handlers.Uhandler.CreateUserhandler)
		r.Get("/user/logout", handlers.Uhandler.LogoutUserhandler)
		r.Get("/user/login", handlers.Uhandler.LoginUserhandler)
		r.Delete("/user/{username}", handlers.Uhandler.DeleteUserhandler)
		r.Put("/user/{username}", handlers.Uhandler.UpdateUserhandler)
		r.Get("/user/{username}", handlers.Uhandler.GetUserhandler)
		r.Post("/user/createWithList", handlers.Uhandler.CreateUserWithArrayhandler)
		r.Post("/user/createWithArray", handlers.Uhandler.CreateUserWithArrayhandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token))
		r.Use(handlers.Uhandler.Authenticator)

		r.Get("/store/inventory", handlers.SHandler.GetInventoryOrderHandler)
	})

	r.Group(func(r chi.Router) {
		r.Delete("/store/order/{orderId}", handlers.SHandler.DeleteOrderHandler)
		r.Get("/store/order/{orderId}", handlers.SHandler.FindOrderHandler)
		r.Post("/store/order", handlers.SHandler.CreateOrderHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token))
		r.Use(handlers.Uhandler.Authenticator)

		r.Post("/pet", handlers.PHandler.CreatePetHandler)
		r.Put("/pet", handlers.PHandler.PutPetHandler)
		r.Get("/pet/findByStatus", handlers.PHandler.FindByStatusPetHandler)
		r.Get("/pet/{petId}", handlers.PHandler.FindByIdPetHandler)
		r.Post("/pet/{petId}", handlers.PHandler.UpdateByIdPetHandler)
		r.Delete("/pet/{petId}", handlers.PHandler.DeleteByIdHandler)
	})

	return r
}

func SwaggerHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/docs") {
			http.ServeFile(w, r, "/docs/swagger.json")
		} else if strings.HasPrefix(r.URL.String(), "/swagger") {
			swag.SwaggerUI(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
