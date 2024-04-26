package main

import (
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginRequest User

	RegisterRequest User

	RegisterResponse struct {
		Status string `json:"status"`
	}

	LoginResponse struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}
)

var Users = map[string]User{}

// @Summary		Login
// @Tags			auth
// @Description	login account
// @ID				login
// @Accept			json
// @Produce		json
// @Param			input	body		LoginRequest	true	"request"
// @Success		200		{object}	LoginResponse
// @Failure		403		{string}	string	"invalid token"
// @Router			/api/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody LoginRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Неверный формат запроса", 400)
		return
	}

	var tokenString string
	if val, ok := Users[reqBody.Username]; ok {
		err = bcrypt.CompareHashAndPassword([]byte(val.Password), []byte(reqBody.Password))
		if err != nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("No such passwd"))
			return
		}
		_, tokenString, _ = tokenAuth.Encode(map[string]interface{}{"username": val.Username})
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("No such user"))
		return
	}

	respBody := LoginResponse{
		Status: http.StatusText(http.StatusOK),
		Token:  tokenString,
	}

	body, err = json.Marshal(respBody)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

// @Summary		Register
// @Tags			auth
// @Description	Post Register
// @ID				create-account
// @Accept			json
// @Produce		json
// @Param			input	body		RegisterRequest	true	"request"
// @Success		200		{object}	RegisterResponse
// @Failure		403		{string}	string	"invalid token"
// @Router			/api/register [post]
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody RegisterRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Неверный формат запроса", 400)
		return
	}
	Users[reqBody.Username] = User{
		Username: reqBody.Username,
		Password: hashPasswd(reqBody.Password),
	}

	respBody := RegisterResponse{http.StatusText(http.StatusOK)}
	body, err = json.Marshal(respBody)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func hashPasswd(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(hash)
}
