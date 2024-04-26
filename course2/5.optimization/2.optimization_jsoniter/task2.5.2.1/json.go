package main

import (
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
)

const (
	System CreatedBy = "system"
)

type Welcome struct {
	Records     []Record `json:"records"`
	Skip        int64    `json:"skip"`
	Limit       int64    `json:"limit"`
	TotalAmount int64    `json:"totalAmount"`
}

type Record struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Profile   Profile   `json:"profile"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	CreatedAt string    `json:"createdAt"`
	CreatedBy CreatedBy `json:"createdBy"`
}

type Profile struct {
	Dob        string `json:"dob"`
	Avatar     string `json:"avatar"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	StaticData string `json:"staticData"`
}

type CreatedBy string

type MarshalUnmarshaler interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

type CheckJson struct {
}

type CheckJsoniter struct {
}

type CheckEasyjson struct {
}

func (c *CheckJson) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
func (c *CheckJson) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (c *CheckJsoniter) Marshal(v any) ([]byte, error) {
	return jsoniter.Marshal(v)
}
func (c *CheckJsoniter) Unmarshal(data []byte, v any) error {
	return jsoniter.Unmarshal(data, v)
}

func (c *CheckEasyjson) Marshal(v any) ([]byte, error) {
	return easyjson.Marshal(v.(easyjson.Marshaler))
}
func (c *CheckEasyjson) Unmarshal(data []byte, v any) error {
	return easyjson.Unmarshal(data, v.(easyjson.Unmarshaler))
}
