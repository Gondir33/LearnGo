package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"repository/internal/infrastructure/db/dao"
	"repository/internal/infrastructure/responder"
	"repository/internal/storage"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ptflp/godecoder"
	"go.uber.org/zap"
)

func TestSearchHandler(t *testing.T) {
	// Create a sample request body
	reqBody := DeleteUserRequest{
		Id: "5",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	// Create a request with the sample body
	req, err := http.NewRequest("POST", "/api/users/delete", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	logg, _ := zap.NewProduction()
	dbRaw, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "db", "5432", "postgres", "yourpassword", "postgres"))
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDb(dbRaw, "postgres")
	sqlAdapter := dao.NewDAO(dbx)
	user := NewUserer(storage.NewUserStorage(sqlAdapter), responder.NewResponder(godecoder.NewDecoder(), logg))
	// Call the handler function
	user.DeleteUserhandler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := `["г Москва, ул Сухонская, д 11"]` // Update with expected addresses
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
