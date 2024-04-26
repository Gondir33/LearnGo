package address

/*
func TestSearchHandler(t *testing.T) {
	// Create a sample request body
	reqBody := SearchRequest{
		Query: "москва сухонская 11",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	// Create a request with the sample body
	req, err := http.NewRequest("POST", "/search", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function
	SearchHandler(rr, req)

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

func TestGeocodeHandler(t *testing.T) {
	// Create a sample request body
	reqBody := GeocodeRequest{
		Lat: "55.19",
		Lng: "61.34",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	// Create a request with the sample body
	req, err := http.NewRequest("POST", "/geocode", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function
	GeocodeHandler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := `["Address1", "Address2"]` // Update with expected addresses
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
*/
