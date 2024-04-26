package service

import "proxy/internal/models"

type GeoServicer interface {
	SearchApi(*SearchRequest) (*SearchResponse, error)
	Login(*LoginRequest) (*LoginResponse, error)
	Register(*RegisterRequest) (*RegisterResponse, error)
}

type (
	SearchRequest struct {
		Query string `json:"query"`
	}
	SearchResponse struct {
		Addresses []models.Address `json:"addresses"`
	}

	GeocodeRequest models.Address

	GeocodeResponse struct {
		Addresses []models.Address `json:"addresses"`
	}

	GeocodeJson struct {
		Suggestions []Suggestion `json:"suggestions"`
	}

	Suggestion struct {
		Value             string                 `json:"value"`
		UnrestrictedValue string                 `json:"unrestricted_value"`
		Data              map[string]interface{} `json:"data"`
	}

	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginRequest User

	RegisterRequest User

	RegisterResponse struct {
		Status string `json:"status"`
	}

	LoginResponse struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}
)
