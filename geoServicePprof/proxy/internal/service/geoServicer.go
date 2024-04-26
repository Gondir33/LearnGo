package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

type GeoService struct {
	Users map[string]User
}

func NewGeoService() *GeoService {
	return &GeoService{
		Users: make(map[string]User),
	}
}

func (g *GeoService) SearchApi(SearchR *SearchRequest) (*SearchResponse, error) {
	var respBody SearchResponse

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
		return &respBody, err
	}

	for _, address := range addresses {
		respBody.Addresses = append(respBody.Addresses, &Address{Lat: address.Data.GeoLat, Lng: address.Data.GeoLon})
	}
	return &respBody, nil
}

func (g *GeoService) GeocodeApi(GeocodeR *GeocodeRequest) (*GeocodeResponse, error) {
	var geocodeResp GeocodeResponse

	client := &http.Client{}
	dataBuffer, err := json.Marshal(GeocodeR)
	if err != nil {
		return &geocodeResp, err
	}
	req, err := http.NewRequest("POST", "http://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", bytes.NewBuffer(dataBuffer))
	if err != nil {
		return &geocodeResp, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token 0eade739b4a52041d493615d92913bfa4a2ebcab")

	respBody, err := client.Do(req)
	if err != nil {
		return &geocodeResp, err
	}
	defer respBody.Body.Close()

	bodyText, err := io.ReadAll(respBody.Body)
	if err != nil {
		return &geocodeResp, err
	}
	var dataSuggestions GeocodeJson
	err = json.Unmarshal(bodyText, &dataSuggestions)
	if err != nil {
		return &geocodeResp, err
	}

	for _, data := range dataSuggestions.Suggestions {
		geocodeResp.Addresses = append(geocodeResp.Addresses,
			&Address{Lat: fmt.Sprintf("%v", data.Data["geo_lat"]), Lng: fmt.Sprintf("%v", data.Data["geo_lon"])})
	}
	return &geocodeResp, nil
}

func (g *GeoService) Register(RegisterR *RegisterRequest) (*RegisterResponse, error) {
	g.Users[RegisterR.Username] = User{
		Username: RegisterR.Username,
		Password: hashPasswd(RegisterR.Password),
	}

	respBody := RegisterResponse{http.StatusText(http.StatusOK)}

	return &respBody, nil
}
func hashPasswd(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(hash)
}

func (g *GeoService) Login(LoginR *LoginRequest) (*LoginResponse, error) {
	var respBody LoginResponse

	var tokenString string
	if val, ok := g.Users[LoginR.Username]; ok {
		err := bcrypt.CompareHashAndPassword([]byte(val.Password), []byte(LoginR.Password))
		if err != nil {
			return &respBody, errors.New("No such passwd")
		}
		_, tokenString, _ = TokenAuth.Encode(map[string]interface{}{"username": val.Username})
	} else {
		return &respBody, errors.New("No such user")
	}

	respBody = LoginResponse{
		Status: http.StatusText(http.StatusOK),
		Token:  tokenString,
	}

	return &respBody, nil
}
