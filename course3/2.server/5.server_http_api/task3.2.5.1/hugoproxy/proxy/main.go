package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"test/address"

	_ "test/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//	@title			HugoMap
//	@version		1.0
//	@description	API Server for HugoMap Application
//
//	@host			localhost:8080
//	@BasePath		/
func main() {
	r := chi.NewRouter()

	rp := NewReverseProxy("hugo", "1313")

	r.Use(middleware.Logger)
	r.Use(rp.ReverseProxy)

	r.Post("/api/address/search", address.SearchHandler)
	r.Post("/api/address/geocode", address.GeocodeHandler)

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

/*
func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	reverseProxyURL, err := url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
	if err != nil {
		log.Fatal(err)
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(reverseProxyURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/docs") {
			http.ServeFile(w, r, "/docs/swagger.json")
			return
		}
		if strings.HasPrefix(r.URL.Path, "/swagger") {
			swaggerUI(w, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/api") {
			next.ServeHTTP(w, r)
		}
		reverseProxy.ServeHTTP(w, r)
	})
}
*/
