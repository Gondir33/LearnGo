package controller

import (
	"golibrary/internal/infrastructure/responder"
	"golibrary/internal/models"
	"golibrary/internal/modules/user/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/ptflp/godecoder"
)

type Userer interface {
	CreateUserhandler(http.ResponseWriter, *http.Request)
	GetUserByUsernamehandler(http.ResponseWriter, *http.Request)
	GetUserByIdhandler(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
}

type User struct {
	user      service.Userer
	Responder responder.Responder
	Decoder   godecoder.Decoder
}

func NewUserHandler(user service.Userer, respond responder.Responder, Decoder godecoder.Decoder) Userer {
	return &User{
		user:      user,
		Responder: respond,
		Decoder:   Decoder,
	}
}

// @Summary	Create user
// @Tags		user
// @Accept		json
// @Produce	json
// @Param		request	body	CreateUserRequest	true	"Account info"
// @Success	200
// @Router		/user [post]
func (u *User) CreateUserhandler(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateUserRequest

	err := u.Decoder.Decode(r.Body, &reqBody)
	if err != nil {
		u.Responder.ErrorBadRequest(w, err)
		return
	}
	err = u.user.Create(r.Context(), models.User(reqBody))
	if err != nil {
		u.Responder.ErrorInternal(w, err)
		return
	}
	u.Responder.OutputJSON(w, http.StatusOK)
}

// @Summary	Get user by username
// @Tags		user
// @Accept		json
// @Produce	json
// @Param		username	path		string	true	"username"
// @Success	200			{object}	models.UserDTO
// @Router		/user/{username} [get]
func (u *User) GetUserByUsernamehandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := u.user.GetUserByUsername(r.Context(), username)
	if err != nil {
		u.Responder.ErrorInternal(w, err)
		return
	}
	u.Responder.OutputJSON(w, user)
}

// @Summary	Get user by id
// @Tags		user
// @Accept		json
// @Produce	json
// @Param		id	path		int	true	"id"
// @Success	200	{object}	models.UserDTO
// @Router		/user/{id} [get]
func (u *User) GetUserByIdhandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		u.Responder.ErrorBadRequest(w, err)
		return
	}
	user, err := u.user.GetUserById(r.Context(), idInt)
	if err != nil {
		u.Responder.ErrorInternal(w, err)
		return
	}
	u.Responder.OutputJSON(w, user)
}

// @Summary	List Users
// @Tags		user
// @Accept		json
// @Produce	json
// @Success	200			{object}	[]models.UserDTO
// @Router		/user/list [get]
func (u *User) List(w http.ResponseWriter, r *http.Request) {
	users, err := u.user.List(r.Context())
	if err != nil {
		u.Responder.ErrorInternal(w, err)
		return
	}
	u.Responder.OutputJSON(w, users)
}

type CreateUserRequest models.User
