package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"proxy/internal/service"
	"proxy/metrics"
	"time"
)

// @Summary		Register
// @Tags			auth
// @Description	Post Register
// @ID				create-account
// @Accept			json
// @Produce		json
// @Param			input	body		service.RegisterRequest	true	"request"
// @Success		200		{object}	service.RegisterResponse
// @Failure		403		{string}	string	"invalid token"
// @Router			/api/register [post]
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody service.RegisterRequest

	start := time.Now()
	defer func() {
		metrics.ObserveRequsetDuration(time.Since(start), "Register")
		metrics.ObserveRequsetCount("Register")
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respond.ErrorInternal(w, errors.New("Ошибка чтения тела запроса"))
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		respond.ErrorBadRequest(w, errors.New("Неверный формат запроса"))
		return
	}

	respBody, err := geo.Register(&reqBody)
	if err != nil {
		respond.OutputJSON(w, err.Error())
		return
	}
	respond.OutputJSON(w, respBody)
}

// @Summary		Login
// @Tags			auth
// @Description	login account
// @ID				login
// @Accept			json
// @Produce		json
// @Param			input	body		service.LoginRequest	true	"request"
// @Success		200		{object}	service.LoginResponse
// @Failure		403		{string}	string	"invalid token"
// @Router			/api/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody service.LoginRequest

	start := time.Now()
	defer func() {
		metrics.ObserveRequsetDuration(time.Since(start), "Login")
		metrics.ObserveRequsetCount("Login")
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respond.ErrorInternal(w, errors.New("Ошибка чтения тела запроса"))
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		respond.ErrorBadRequest(w, errors.New("Неверный формат запроса"))
		return
	}

	respBody, err := geo.Login(&reqBody)
	if err != nil {
		respond.OutputJSON(w, err.Error())
		return
	}
	respond.OutputJSON(w, respBody)
}
