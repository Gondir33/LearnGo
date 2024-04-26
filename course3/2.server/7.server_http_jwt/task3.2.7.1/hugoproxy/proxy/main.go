package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"proxy/address"
	"strings"

	_ "proxy/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

//	@title			HugoMap
//	@version		1.0
//	@description	API Server for HugoMap Application
//
//	@host			localhost:8080
//	@BasePath		/

//	@securityDefinitions.apikey	ApiKeyAuth
//
// @in Login
// @name Authorization
func main() {
	r := chi.NewRouter()

	rp := NewReverseProxy("hugo", "1313")
	r.Use(middleware.Logger)
	r.Use(rp.ReverseProxy)

	r.Group(func(r chi.Router) {
		r.Post("/api/login", LoginHandler)
		r.Post("/api/register", RegisterHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/api/address/search", address.SearchHandler)
		r.Post("/api/address/geocode", address.GeocodeHandler)
	})

	r.Get("/", defaultHandler)

	http.ListenAndServe(":8080", r)
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/docs") {
			http.ServeFile(w, r, "/docs/swagger.json")
		} else if strings.HasPrefix(r.URL.String(), "/swagger") {
			swaggerUI(w, r)
		} else if strings.HasPrefix(r.URL.String(), "/api") {
			next.ServeHTTP(w, r)
		} else {
			url, _ := url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
			proxy := httputil.NewSingleHostReverseProxy(url)
			proxy.ServeHTTP(w, r)
		}
	})
}
func defaultHandler(w http.ResponseWriter, r *http.Request) {
}
