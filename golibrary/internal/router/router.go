package router

import (
	"golibrary/internal/infrastructure/component"
	"golibrary/internal/modules"
	"net/http"

	"github.com/go-chi/chi"
)

func NewApiRouter(controllers *modules.Controllers, components *component.Components) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/user", controllers.Userer.CreateUserhandler)
		r.Get("/user/{username}", controllers.Userer.GetUserByUsernamehandler)
		r.Get("/user/{id} ", controllers.Userer.GetUserByIdhandler)
		r.Get("/user/list", controllers.Userer.List)
	})
	r.Group(func(r chi.Router) {
		r.Post("/book/{id_author}", controllers.Libraryer.CreateBook)
		r.Post("/author", controllers.Libraryer.CreateAuthor)
		r.Get("/list/book", controllers.Libraryer.ListBook)
		r.Get("/list/authors", controllers.Libraryer.ListAuthor)
		r.Post("/book/get", controllers.Libraryer.GetBook)
		r.Put("/book/hand", controllers.Libraryer.HandBook)
		r.Get("/rentedbook", controllers.Libraryer.RentedBooks)
	})
	return r
}
