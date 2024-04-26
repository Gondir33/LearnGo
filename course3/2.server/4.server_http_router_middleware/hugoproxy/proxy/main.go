package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"test/workers"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	rp := NewReverseProxy("hugo", "1313")

	r.Use(rp.ReverseProxy)

	r.HandleFunc("/", defaultHandler)

	workers.Workers()

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
		if strings.Contains(r.URL.String(), "api") {
			fmt.Fprint(w, "Hello from API")
		} else {
			url, _ := url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
			proxy := httputil.NewSingleHostReverseProxy(url)
			proxy.ServeHTTP(w, r)
		}
	})
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
}
