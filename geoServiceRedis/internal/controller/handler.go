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
// @Param			input	body		service.SearchRequest	true	"request"
// @Success		200		{object}	service.SearchResponse
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
