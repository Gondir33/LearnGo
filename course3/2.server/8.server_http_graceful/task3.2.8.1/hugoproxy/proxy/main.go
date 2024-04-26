package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"proxy/address"
	"strings"
	"syscall"
	"time"

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
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil, // Здесь должен быть ваш обработчик запросов
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting server...")
		if err := Server(server); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Ожидание сигнала остановки
	<-stopChan

	// Создание контекста с таймаутом в пять секунд для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Остановка сервера с использованием graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped gracefully")
}

func Server(server *http.Server) error {
	r := chi.NewRouter()
	server.Handler = r

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

	return server.ListenAndServe()
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
