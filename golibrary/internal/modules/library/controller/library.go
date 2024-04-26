package controller

import (
	"errors"
	"golibrary/internal/infrastructure/responder"
	"golibrary/internal/models"
	"golibrary/internal/modules/library/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/ptflp/godecoder"
)

type Libraryer interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
	CreateAuthor(w http.ResponseWriter, r *http.Request)
	ListBook(w http.ResponseWriter, r *http.Request)
	ListAuthor(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	HandBook(w http.ResponseWriter, r *http.Request)
	RentedBooks(w http.ResponseWriter, r *http.Request)
}

type Library struct {
	Library   service.LibraryServicere
	Responder responder.Responder
	Decoder   godecoder.Decoder
}

func NewLibraryHandler(LibraryServicere service.LibraryServicere, respond responder.Responder, Decoder godecoder.Decoder) Libraryer {
	return &Library{
		Library:   LibraryServicere,
		Responder: respond,
		Decoder:   Decoder,
	}
}

// @Summary	Create book
// @Tags		library
// @Accept		json
// @Produce	json
// @Param		request		body	models.Book	true	"Book info"
// @Param		id_author	path	int			true	"id_author"
// @Success	200
// @Router		/book/{id_author} [post]
func (l *Library) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	id_author := chi.URLParam(r, "id_author")
	err := l.Decoder.Decode(r.Body, &book)
	if err != nil {
		l.Responder.ErrorBadRequest(w, err)
		return
	}
	id, err := strconv.Atoi(id_author)
	if err != nil {
		l.Responder.ErrorBadRequest(w, err)
		return
	}
	if err = l.Library.CreateBook(r.Context(), book, id); err != nil {
		l.Responder.ErrorInternal(w, err)
		return
	}
	l.Responder.OutputJSON(w, http.StatusOK)
}

// @Summary	Create Author
// @Tags		library
// @Accept		json
// @Produce	json
// @Param		request	body	models.Author	true	"Account info"
// @Success	200
// @Router		/author [post]
func (l *Library) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	err := l.Decoder.Decode(r.Body, &author)
	if err != nil {
		l.Responder.ErrorBadRequest(w, err)
		return
	}
	if err = l.Library.CreateAuthor(r.Context(), author); err != nil {
		l.Responder.ErrorInternal(w, err)
		return
	}
	l.Responder.OutputJSON(w, http.StatusOK)
}

// @Summary	Get Books
// @Tags		library
// @Accept		json
// @Produce	json
// @Success	200	{object}	[]models.BookWithAuthor
// @Router		/list/book [get]
func (l *Library) ListBook(w http.ResponseWriter, r *http.Request) {
	books, err := l.Library.ListBook(r.Context())
	if err != nil {
		l.Responder.ErrorInternal(w, err)
		return
	}
	l.Responder.OutputJSON(w, books)
}

// @Summary	Get Authors
// @Tags		library
// @Accept		json
// @Produce	json
// @Success	200	{object}	[]models.AuthorWithBooks
// @Router		/list/authors [get]
func (l *Library) ListAuthor(w http.ResponseWriter, r *http.Request) {
	books, err := l.Library.ListAuthor(r.Context())
	if err != nil {
		l.Responder.ErrorInternal(w, err)
		return
	}
	l.Responder.OutputJSON(w, books)
}

// @Summary	get id_book for id_user
// @Tags		library
// @Accept		json
// @Produce	json
// @Param		id_book	query	int	true	"id_book"
// @Param		id_user	query	int	true	"id_user"
// @Success	200
// @Router		/book/get [post]
func (l *Library) GetBook(w http.ResponseWriter, r *http.Request) {
	id_book := r.URL.Query().Get("id_book")
	id_user := r.URL.Query().Get("id_user")
	if id_book == "" || id_user == "" {
		l.Responder.ErrorBadRequest(w, errors.New("no id_book or id_user"))
		return
	}

	idBook, _ := strconv.Atoi(id_book)
	idUser, err := strconv.Atoi(id_user)
	if err != nil {
		l.Responder.ErrorBadRequest(w, err)
		return
	}

	err = l.Library.GetBook(r.Context(), idUser, idBook)
	if err != nil {
		l.Responder.ErrorInternal(w, err)
		return
	}
	l.Responder.OutputJSON(w, http.StatusOK)
}

// @Summary	hand id_book for id_user
// @Tags		library
// @Accept		json
// @Produce	json
// @Param		id_book	query	int	true	"id_book"
// @Param		id_user	query	int	true	"id_user"
// @Success	200
// @Router		/book/hand [put]
func (l *Library) HandBook(w http.ResponseWriter, r *http.Request) {
	id_book := r.URL.Query().Get("id_book")
	id_user := r.URL.Query().Get("id_user")
	if id_book == "" || id_user == "" {
		l.Responder.ErrorBadRequest(w, errors.New("no id_book or id_user"))
		return
	}

	idBook, _ := strconv.Atoi(id_book)
	idUser, err := strconv.Atoi(id_user)
	if err != nil {
		l.Responder.ErrorBadRequest(w, err)
		return
	}

	err = l.Library.HandBook(r.Context(), idUser, idBook)
	if err != nil {
		l.Responder.ErrorInternal(w, err)
		return
	}
	l.Responder.OutputJSON(w, http.StatusOK)
}

// @Summary	RentedBooks
// @Tags		library
// @Accept		json
// @Produce	json
// @Success	200	{object}	[]models.UsersWithBook
// @Router		/rentedbook [get]
func (l *Library) RentedBooks(w http.ResponseWriter, r *http.Request) {
	list, err := l.Library.RentedBooks(r.Context())
	if err != nil {
		l.Responder.ErrorInternal(w, err)
		return
	}
	l.Responder.OutputJSON(w, list)
}
