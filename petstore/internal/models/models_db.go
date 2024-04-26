package models

type UserDB struct {
	Id         int    `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
	Username   string `json:"username"  db:"username" db_type:"VARCHAR(100)"`
	FirstName  string `json:"first_name" db:"firstname" db_type:"VARCHAR(100)"`
	LastName   string `json:"last_name" db:"lastname" db_type:"VARCHAR(100)"`
	Email      string `json:"email" db:"email" db_type:"VARCHAR(100)"`
	Password   string `json:"password" db:"password" db_type:"VARCHAR(100)"`
	Phone      string `json:"phone" db:"phone" db_type:"VARCHAR(100)"`
	UserStatus int    `json:"user_status" db:"userstatus" db_type:"INTEGER"`
	CreatedAt  string `db:"created_at" db_type:"VARCHAR(100)"`
	DeletedAt  string `db:"deleted_at" db_type:"VARCHAR(100)"`

	// User Status Schema(description = "User Status", allowableValues = "1-registered,2-active,3-closed)
}

func (user *UserDB) TableName() string {
	return "users"
}

type PetDB struct {
	Id        int    `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
	Category  string `json:"category" db:"category" db_type:"VARCHAR(65535)"`
	Name      string `json:"name" db:"name" db_type:"VARCHAR(100)"`
	PhotoUrls string `json:"photo_urls" db:"photourls" db_type:"VARCHAR(65535)"`
	Tags      string `json:"tags" db:"tags" db_type:"VARCHAR(65535)"`
	Status    string `json:"status" db:"status" db_type:"VARCHAR(100)"`
	CreatedAt string `db:"created_at" db_type:"VARCHAR(100)"`
	DeletedAt string `db:"deleted_at" db_type:"VARCHAR(100)"`
	// Petstatus  Schema(description = "User Status", allowableValues = "1-available,2-pending,3-sold)
}

func (pets PetDB) TableName() string {
	return "pets"
}

type OrderDB struct {
	Id        int    `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
	PetId     int    `json:"pet_id" db:"petid" db_type:"INTEGER"`
	Quantity  int    `json:"quantity" db:"quantity" db_type:"INTEGER"`
	ShipDate  string `json:"ship_date" db:"shipdate" db_type:"VARCHAR(100)"`
	Status    string `json:"status" db:"status" db_type:"VARCHAR(100)"`
	Complete  bool   `json:"complete" db:"complete" db_type:"BOOLEAN NOT NULL"`
	CreatedAt string `db:"created_at" db_type:"VARCHAR(100)"`
	DeletedAt string `db:"deleted_at" db_type:"VARCHAR(100)"`
	// Order  Schema(description = "User Status", allowableValues = "1-placed,2-approved,3-delivered)
}

func (order OrderDB) TableName() string {
	return "orders"
}

type Api_keyDB struct {
	Api_key string `json:"api_key" db:"api_key" db_type:"VARCHAR(100)"`
}

func (a Api_keyDB) TableName() string {
	return "api_keys"
}
