package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_server(t *testing.T) {
	tests := []struct {
		name            string
		arg             string
		expectedMessage string
		expectedStatus  int
		handler         http.HandlerFunc
	}{{
		name:            "/1",
		arg:             "/1",
		expectedMessage: "Hello world",
		expectedStatus:  http.StatusOK,
		handler:         firstHandler,
	}, {
		name:            "/2",
		arg:             "/2",
		expectedMessage: "Hello world 2",
		expectedStatus:  http.StatusOK,
		handler:         secondHandler,
	}, {
		name:            "/3",
		arg:             "/3",
		expectedMessage: "Hello world 3",
		expectedStatus:  http.StatusOK,
		handler:         thirdHandler,
	}, {
		name:            "/",
		arg:             "/asdasdsa/asdsa/",
		expectedMessage: "Not Found",
		expectedStatus:  http.StatusNotFound,
		handler:         defaultHandler,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.arg, nil)
			responseRecorder := httptest.NewRecorder()
			tt.handler(responseRecorder, request)
			if responseRecorder.Code != tt.expectedStatus {
				t.Errorf("Want status '%v', got '%v'", tt.expectedStatus, responseRecorder.Code)
			}
			if responseRecorder.Body.String() != tt.expectedMessage {
				t.Errorf("Want message '%v', got '%v'", tt.expectedMessage, responseRecorder.Body.String())
			}
		})
	}
}
