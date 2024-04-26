package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"repository/internal/infrastructure/db/dao"
	"repository/internal/infrastructure/responder"
	"repository/internal/models"
	"repository/internal/storage"

	"github.com/go-chi/chi"
)

type Userer interface {
	CreateUserhandler(w http.ResponseWriter, r *http.Request)
	UpdateUserhandler(w http.ResponseWriter, r *http.Request)
	GetByIdUserhandler(w http.ResponseWriter, r *http.Request)
	DeleteUserhandler(w http.ResponseWriter, r *http.Request)
	ListUserhandler(w http.ResponseWriter, r *http.Request)
}

type User struct {
	userRep storage.UserRepository
	respond responder.Responder
}

func NewUserer(userRep storage.UserRepository, respond responder.Responder) Userer {
	return &User{
		userRep: userRep,
		respond: respond,
	}
}

// @Summary		Create
// @Tags			user
// @Description	Post Create
// @Accept			json
// @Produce		json
// @Param			input	body		CreateUserRequest	true	"request"
// @Success		200		{object}	int
// @Failure		403		{string}	error
// @Router			/api/users/create [post]
func (u *User) CreateUserhandler(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateUserRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		u.respond.ErrorBadRequest(w, err)
		return
	}

	err = u.userRep.Create(context.Background(), models.User{
		Id:       reqBody.Id,
		Username: reqBody.Username,
		Password: reqBody.Password,
	})
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	u.respond.OutputJSON(w, http.StatusOK)
}

// @Summary		Update
// @Tags			user
// @Description	Post Update
// @Accept			json
// @Produce		json
// @Param			input	body		UpdateUserRequest	true	"request"
// @Success		200		{object}	int
// @Failure		403		{string}	error
// @Router			/api/users/update [post]
func (u *User) UpdateUserhandler(w http.ResponseWriter, r *http.Request) {
	var reqBody UpdateUserRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		u.respond.ErrorBadRequest(w, err)
		return
	}

	err = u.userRep.Update(context.Background(), models.User{
		Id:       reqBody.Id,
		Username: reqBody.Username,
		Password: reqBody.Password,
	})
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	u.respond.OutputJSON(w, http.StatusOK)
}

// @Summary		GetByID
// @Tags			user
// @Description	Get By id
// @Accept			json
// @Produce		json
// @Param			input	body		string false	"request"
// @Success		200		{object}	GetByIdUserResponse
// @Failure		403		{string}	error
// @Router			/api/users/{id} [get]
func (u *User) GetByIdUserhandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := u.userRep.GetByID(context.Background(), id)
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}

	u.respond.OutputJSON(w, GetByIdUserResponse(user))
}

// @Summary		Delete
// @Tags			user
// @Description	Post Delet
// @Accept			json
// @Produce		json
// @Param			input	body		DeleteUserRequest true	"request"
// @Success		200		{object}	int
// @Failure		403		{string}	error
// @Router			/api/users/delete [post]
func (u *User) DeleteUserhandler(w http.ResponseWriter, r *http.Request) {
	var reqBody DeleteUserRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		u.respond.ErrorBadRequest(w, err)
		return
	}

	err = u.userRep.Delete(context.Background(), reqBody.Id)
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	u.respond.OutputJSON(w, http.StatusOK)
}

// @Summary		List
// @Tags			user
// @Description	Post List
// @Accept			json
// @Produce		json
// @Param			input	body		ListUserRequest true	"request"
// @Success		200		{object}	ListUserResponse
// @Failure		403		{string}	error
// @Router			/api/users/list [post]
func (u *User) ListUserhandler(w http.ResponseWriter, r *http.Request) {
	var reqBody ListUserRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		u.respond.ErrorBadRequest(w, err)
		return
	}

	users, err := u.userRep.List(context.Background(), dao.Condition{
		LimitOffset: &dao.LimitOffset{
			Offset: reqBody.Offset,
			Limit:  reqBody.Limit,
		},
	})
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	u.respond.OutputJSON(w, ListUserResponse(users))
}

type (
	CreateUserRequest struct {
		Id       int    `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
		Username string `json:"username" db:"username" db_type:"VARCHAR(100)"`
		Password string `json:"password" db:"password" db_type:"VARCHAR(100)"`
	}
	UpdateUserRequest struct {
		Id       int    `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
		Username string `json:"username" db:"username" db_type:"VARCHAR(100)"`
		Password string `json:"password" db:"password" db_type:"VARCHAR(100)"`
	}
	GetByIdUserResponse models.User
	DeleteUserRequest   struct {
		Id string `json:"id"`
	}
	ListUserRequest  dao.LimitOffset
	ListUserResponse []models.User
)
