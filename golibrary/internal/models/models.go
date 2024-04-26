package models

type Book struct {
	Name string `json:"name"`
}

type Author struct {
	Name string `json:"name"`
}

type User struct {
	Name string `json:"name"`
}

type BookDTO struct {
	Id   int
	Name string
}

type AuthorDTO struct {
	Id   int
	Name string
}

type UserDTO struct {
	Id   int
	Name string
}

type AuthorWithBooks struct {
	Author Author
	Books  []Book
}
type BookWithAuthor struct {
	Author Author
	Book   Book
}

type UsersWithBook struct {
	User        User
	RentedBooks []BookWithAuthor
}
