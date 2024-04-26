package address

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "test/docs"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
)

// @Summary		Search
// @Description	Post Address
// @Tags			search
// @Accept			json
// @Produce		json
// @Param			input	body		SearchRequest	true	"request"
// @Success		200		{object}		SearchResponse
// @Failure		400		{string}	string	"Неверный формат запроса"
// @Failure		404		{string}	string	"404 not found"
// @Failure		500		{string}	string	"Cервис https://dadata.ru не доступен"
// @Router			/api/address/search [post]
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody SearchRequest

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

	resp, err := SearchApi(&reqBody, w, r)
	if err != nil {
		return
	}
	ans, err := json.Marshal(resp.Addresses)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(ans)
}

func SearchApi(SearchR *SearchRequest, w http.ResponseWriter, r *http.Request) (*SearchResponse, error) {
	var resp SearchResponse

	creds := client.Credentials{
		ApiKeyValue:    "0eade739b4a52041d493615d92913bfa4a2ebcab",
		SecretKeyValue: "2a6a465a6d34c9ad838cf089cb3cc323ab2fa296",
	}
	api := dadata.NewSuggestApi(client.WithCredentialProvider(&creds))

	params := suggest.RequestParams{
		Query: SearchR.Query,
	}

	addresses, err := api.Address(context.Background(), &params)
	if err != nil {
		http.Error(w, "Cервис https://dadata.ru не доступен", http.StatusInternalServerError)
		return &resp, err
	}

	for _, address := range addresses {
		resp.Addresses = append(resp.Addresses, &Address{Lat: address.Data.GeoLat, Lng: address.Data.GeoLon})
	}
	resp.Addresses = resp.Addresses[:len(resp.Addresses)-1]
	return &resp, nil
}

// @Summary		Search
// @Description	Post Address
// @Tags			Geocode
// @Accept			json
// @Produce		json
// @Param			input	body		GeocodeRequest	true	"request"
// @Success		200		{object}		GeocodeResponse
// @Failure		400		{string}	string	"Неверный формат запроса"
// @Failure		404		{string}	string	"404 not found"
// @Failure		500		{string}	string	"Cервис https://dadata.ru не доступен"
// @Router			/api/address/geocode [post]
func GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody GeocodeRequest

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

	resp, err := GeocodeApi(&reqBody, w, r)
	if err != nil {
		return
	}
	ans, err := json.Marshal(resp.Addresses)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(ans)
}

func GeocodeApi(GeocodeR *GeocodeRequest, w http.ResponseWriter, r *http.Request) (*GeocodeResponse, error) {
	var geocodeResp GeocodeResponse

	client := &http.Client{}
	dataBuffer, err := json.Marshal(GeocodeR)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", bytes.NewBuffer(dataBuffer))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token 0eade739b4a52041d493615d92913bfa4a2ebcab")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Cервис https://dadata.ru не доступен", http.StatusInternalServerError)
		return &geocodeResp, err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Cервис https://dadata.ru не доступен", http.StatusInternalServerError)
		return &geocodeResp, err
	}
	var dataSuggestions GeocodeJson
	err = json.Unmarshal(bodyText, &dataSuggestions)
	if err != nil {
		log.Fatal(err)
	}

	for _, data := range dataSuggestions.Suggestions {
		geocodeResp.Addresses = append(geocodeResp.Addresses,
			&Address{Lat: fmt.Sprintf("%v", data.Data["geo_lat"]), Lng: fmt.Sprintf("%v", data.Data["geo_lon"])})
	}
	return &geocodeResp, nil
}

type (
	Address struct {
		Lat string `json:"lat"`
		Lng string `json:"lon"`
	}

	SearchRequest struct {
		Query string `json:"query"`
	}
	SearchResponse struct {
		Addresses []*Address `json:"addresses"`
	}

	GeocodeRequest Address

	GeocodeResponse struct {
		Addresses []*Address `json:"addresses"`
	}

	GeocodeJson struct {
		Suggestions []Suggestion `json:"suggestions"`
	}

	Suggestion struct {
		Value             string                 `json:"value"`
		UnrestrictedValue string                 `json:"unrestricted_value"`
		Data              map[string]interface{} `json:"data"`
	}
)
