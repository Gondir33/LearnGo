// main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func TestGroupHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		statusCode int
		response   string
	}{
		{"ValidParams", "/1", http.StatusOK, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := chi.NewRouter()
			logger, _ := zap.NewProduction()
			logger.Info("The info about clinet")

			router.Use(LoggerMiddleware)

			router.HandleFunc("/1", handleRoute1)

			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.statusCode {
				t.Errorf("Handler returned wrong status code: got %v, want %v", status, tt.statusCode)
			}

			if tt.statusCode == http.StatusOK {
				expected := tt.response
				actual := rr.Body.String()

				if actual != expected {
					t.Errorf("Handler returned unexpected body: got %v, want %v", actual, expected)
				}
			}
		})
	}
}
