package main

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func WithUsername(username string) UserOption {
	return func(u *User) {
		u.Username = username
	}
}

func WithEmail(email string) UserOption {
	return func(u *User) {
		u.Email = email
	}
}

func WithRole(role string) UserOption {
	return func(u *User) {
		u.Role = role
	}
}

func NewUser(id int, options ...UserOption) *User {
	u := &User{ID: id}

	for _, option := range options {
		option(u)
	}
	return u
}

/*
func main() {
	user := NewUser(1, WithUsername("testuser"), WithEmail("testuser@example.com"), WithRole("admin"))
	fmt.Printf("User: %+v\n", user)
}
*/
