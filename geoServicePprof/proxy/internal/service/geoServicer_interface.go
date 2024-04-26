package service

type GeoServicer interface {
	SearchApi(*SearchRequest) (*SearchResponse, error)
	GeocodeApi(*GeocodeRequest) (*GeocodeResponse, error)
	Login(*LoginRequest) (*LoginResponse, error)
	Register(*RegisterRequest) (*RegisterResponse, error)
}

type (
	Address struct {
		Lat string `json:"lat"`
		Lng string `json:"lon"`
	}

	SearchRequest struct {
		Query string `json:"query"`
	}
	SearchResponse struct {
		Addresses []*Address `json:"addresses"`
	}

	GeocodeRequest Address

	GeocodeResponse struct {
		Addresses []*Address `json:"addresses"`
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
