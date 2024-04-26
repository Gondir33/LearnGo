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
	"proxy/internal/controller"
	"proxy/internal/service"
	"strings"
	"syscall"
	"time"

	"net/http/pprof"
	_ "proxy/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/ptflp/godecoder"
	"go.uber.org/zap"
)

//	@title			HugoMap
//	@version		1.0
//	@description	API Server for HugoMap Application
//
//	@host			localhost:8080
//	@BasePath		/

// @securityDefinitions.apikey	ApiKeyAuth
//
// @in							Login
// @name						Authorization
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

	go func() {
		log.Println(PprofServer())
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
		r.Post("/api/login", controller.LoginHandler)
		r.Post("/api/register", controller.RegisterHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(service.TokenAuth))
		r.Use(Authenticator)

		r.Post("/api/address/search", controller.SearchHandler)
		r.Post("/api/address/geocode", controller.GeocodeHandler)
	})

	r.Get("/", defaultHandler)

	return server.ListenAndServe()
}

func PprofServer() error {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		// Как использовать без сваггера не очень понятно, но по факту как я понял
		// зашиту выстраиваем этими двумся строчками, поэтому для того чтобы вытащить
		// файлик profile и trace использовал заккоментированными эти 2 строчки
		// как должно быть в идеале не понял
		// r.Use(jwtauth.Verifier(service.TokenAuth))
		// r.Use(Authenticator)

		r.Get("/mycustompath/pprof", pprof.Index)
		r.Get("/mycustompath/pprof/cmdline", pprof.Cmdline)
		r.Get("/mycustompath/pprof/profile", pprof.Profile)
		r.Get("/mycustompath/pprof/symbol", pprof.Symbol)
		r.Get("/mycustompath/pprof/trace", pprof.Trace)
	})

	return http.ListenAndServe(":6060", r)
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

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())
		logger, _ := zap.NewProduction()
		resp := service.NewResponder(godecoder.NewDecoder(), logger)
		if err != nil {
			resp.ErrorUnauthorized(w, err)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			resp.ErrorUnauthorized(w, err)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
