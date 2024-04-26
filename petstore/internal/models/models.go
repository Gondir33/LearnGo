package models

// User Status Schema(description = "User Status", allowableValues = "1-registered,2-active,3-closed)
type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"userStatus"`
}

type Order struct {
	Id       int    `json:"id"`
	PetId    int    `json:"pet_id"`
	Quantity int    `json:"quantity"`
	ShipDate string `json:"ship_date"`
	Status   string `json:"status"`
	Complete bool   `json:"complete"`
	// OrderStatus  Schema(description = "Order Status", allowableValues = "1-placed,2-approved,3-delivered)
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Pet struct {
	Id        int      `json:"id"`
	Category  Category `json:"category"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photo_urls"`
	Tags      []Tag    `json:"tags"`
	Status    string   `json:"status"`
	// Petstatus  Schema(description = "Pet Status", allowableValues = "1-available,2-pending,3-sold)
}
