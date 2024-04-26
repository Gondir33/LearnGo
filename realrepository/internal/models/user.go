package models

type User struct {
	Id        int    `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
	Username  string `json:"username" db:"username" db_type:"VARCHAR(100)"`
	Password  string `json:"password" db:"password" db_type:"VARCHAR(100)"`
	DeletedAt string `db:"deleted_at" db_type:"VARCHAR(100)"`
}

func (user User) TableName() string {
	return "users"
}
