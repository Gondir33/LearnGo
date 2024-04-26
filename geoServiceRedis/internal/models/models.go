package models

import (
	"encoding/json"
)

type Address struct {
	Lat string `json:"lat"`
	Lng string `json:"lon"`
}

func (a Address) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}
