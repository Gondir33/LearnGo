package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"petstore/internal/infrastructure/responder"
	"petstore/internal/models"
	"petstore/internal/modules/user/service"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
)

type Userer interface {
	CreateUserhandler(http.ResponseWriter, *http.Request)
	CreateUserWithArrayhandler(http.ResponseWriter, *http.Request)
	GetUserhandler(http.ResponseWriter, *http.Request)
	UpdateUserhandler(http.ResponseWriter, *http.Request)
	DeleteUserhandler(http.ResponseWriter, *http.Request)
	LoginUserhandler(http.ResponseWriter, *http.Request)
	LogoutUserhandler(http.ResponseWriter, *http.Request)
	Authenticator(next http.Handler) http.Handler
}

type User struct {
	user    service.Userer
	respond responder.Responder
}

func NewUserHandler(user service.Userer, respond responder.Responder) Userer {
	return &User{
		user:    user,
		respond: respond,
	}
}

//	@Summary	Create user
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		request	body	CreateUserRequest	true	"Account info"
//	@Success	200
//	@Router		/user [post]
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
	err = u.user.Create(models.User(reqBody))
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	u.respond.OutputJSON(w, "")
}

//	@Summary	Creates list of users with given input array
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		request	body	CreateUserWithArrayRequest	true	"Account"
//	@Success	200
//	@Router		/user/createWithArray [post]
//	@Router		/user/createWithList [post]
func (u *User) CreateUserWithArrayhandler(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateUserWithArrayRequest

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

	err = u.user.CreateWithArray(reqBody)
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	u.respond.OutputJSON(w, "")
}

//	@Summary	Get user by username
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		username	path		string	true	"username"
//	@Success	200			{object}	models.User
//	@Router		/user/{username} [get]
func (u *User) GetUserhandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := u.user.GetUser(username)
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	u.respond.OutputJSON(w, user)
}

//	@Summary	Update user
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		username	path	string				true	"username"
//	@Param		request		body	UpdateUserRequest	true	"Account"
//	@Success	200
//	@Router		/user/{username} [put]
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
	username := chi.URLParam(r, "username")

	err = u.user.Update(username, models.User(reqBody))
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	u.respond.OutputJSON(w, "")
}

//	@Summary	Delete user
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		username	path	string	true	"username"
//	@Success	200
//	@Router		/user/{username} [delete]
func (u *User) DeleteUserhandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	err := u.user.Delete(username)
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}
	u.respond.OutputJSON(w, "")
}

//	@Summary	logs user into the system
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		username	query		string			true	"username"
//	@Param		password	query		string			true	"password"
//	@Success	200			{string}	string			"Bearer \napi_key"
//	@Header		200			{string}	X-Expires-After	"date in UTC when token expires"
//	@Header		200			{string}	X-Rate-Limit	"calls per hour allowed by the user"
//	@Router		/user/login [get]
func (u *User) LoginUserhandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	token, claims, api_key, err := u.user.Login(username, password)
	if err != nil {
		u.respond.ErrorInternal(w, err)
		return
	}

	w.Header().Set("X-Expires-After", fmt.Sprintf("%v", claims["exp"]))
	w.Header().Set("X-Rate-Limit", fmt.Sprintf("%v", claims["rate_limit"]))

	u.respond.OutputJSON(w, "Bearer "+token)
	u.respond.OutputJSON(w, "api_key "+api_key)
}

//	@Summary	Logs out current logged in user session
//	@Tags		user
//	@Accept		json
//	@Security	ApiKeyAuth
//	@Produce	json
//	@Success	200
//	@Router		/user/logout [get]
func (u *User) LogoutUserhandler(w http.ResponseWriter, r *http.Request) {

	tokenString := jwtauth.TokenFromHeader(r)

	u.user.Logout(tokenString)

	u.respond.OutputJSON(w, "")
}

func (u *User) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			u.respond.ErrorUnauthorized(w, err)
			return
		}
		tokenString := jwtauth.TokenFromHeader(r)
		expirationTime := claims["exp"].(time.Time)
		if time.Now().After(expirationTime) {
			u.respond.ErrorUnauthorized(w, err)
			return
		}
		if u.user.CheckTokenLogout(tokenString) {
			u.respond.ErrorUnauthorized(w, err)
			return
		}
		if err != nil {
			u.respond.ErrorUnauthorized(w, err)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			u.respond.ErrorUnauthorized(w, err)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

type CreateUserRequest models.User
type CreateUserWithArrayRequest []models.User
type UpdateUserRequest models.User
