package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleRoute1(t *testing.T) {
	// Create a request to the "/1" route
	req, err := http.NewRequest("GET", "/1", nil)
	assert.NoError(t, err)

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Call the handler function
	handleRoute1(recorder, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Add more assertions based on the specific behavior of handleRoute1 if needed
}

func TestHandleRoute2(t *testing.T) {
	// Create a request to the "/2" route
	req, err := http.NewRequest("GET", "/2", nil)
	assert.NoError(t, err)

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Call the handler function
	handleRoute2(recorder, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Add more assertions based on the specific behavior of handleRoute2 if needed
}

func TestHandleRoute3(t *testing.T) {
	// Create a request to the "/3" route
	req, err := http.NewRequest("GET", "/3", nil)
	assert.NoError(t, err)

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Call the handler function
	handleRoute3(recorder, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Add more assertions based on the specific behavior of handleRoute3 if needed
}
