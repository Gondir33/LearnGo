package controller

import (
	"encoding/json"
	"io"
	"net/http"

	_ "proxy/docs"
	"proxy/internal/service"

	"github.com/ptflp/godecoder"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction()
	respond   = service.NewResponder(godecoder.NewDecoder(), logger)
	geo       = service.GeoServicer(service.NewGeoService())
)

// @Summary		Search
// @Tags			list
// @Security		ApiKeyAuth
// @Description	Post Address
// @Accept			json
// @Produce		json
// @Param			input	body		SearchRequest	true	"request"
// @Success		200		{object}	SearchResponse
// @Failure		400		{string}	string	"Неверный формат запроса"
// @Failure		404		{string}	string	"404 not found"
// @Failure		500		{string}	string	"Cервис https://dadata.ru не доступен"
// @Router			/api/address/search [post]
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody service.SearchRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respond.ErrorInternal(w, err)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		respond.ErrorBadRequest(w, err)
		return
	}

	respBody, err := geo.SearchApi(&reqBody)
	if err != nil {
		respond.ErrorInternal(w, err)
		return
	}
	respond.OutputJSON(w, respBody.Addresses)
}

// @Summary		Search
// @Security		ApiKeyAuth
// @Tags			list
// @Description	Post Address
// @Accept			json
// @Produce		json
// @Param			input	body		GeocodeRequest	true	"request"
// @Success		200		{object}	GeocodeResponse
// @Failure		400		{string}	string	"Неверный формат запроса"
// @Failure		404		{string}	string	"404 not found"
// @Failure		500		{string}	string	"Cервис https://dadata.ru не доступен"
// @Router			/api/address/geocode [post]
func GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody service.GeocodeRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respond.ErrorInternal(w, err)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		respond.ErrorBadRequest(w, err)
		return
	}

	respBody, err := geo.GeocodeApi(&reqBody)
	if err != nil {
		respond.ErrorInternal(w, err)
		return
	}
	respond.OutputJSON(w, respBody)
}
