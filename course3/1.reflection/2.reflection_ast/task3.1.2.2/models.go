package main

type User struct {
	ID        int
	FirstName *string
	LastName  string
	Username  string
	Email     string
	Address   string
	Status    int
	DeletedAt string
}

type Address struct {
	ID     int
	Street string
	City   string
}
