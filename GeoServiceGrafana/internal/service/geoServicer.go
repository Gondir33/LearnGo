package service

import (
	"context"
	"errors"
	"net/http"
	"proxy/internal/storage"

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

	respBody, err := storage.SomeReposProxy.GetData(context.TODO(), SearchR.Query)
	if err != nil {
		return nil, err
	}
	return &SearchResponse{respBody}, nil
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
