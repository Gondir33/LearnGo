package main

//go:generate grizzly generate main.go
//grizzly:generate
type User struct {
	Id   int
	Name string
	Age  int
}

func getUsersByCondition(users []*User, condition string) []*User {
	newUsers := NewUserCollection(users)
	var youngUsers *UserCollection
	// реализуйте функцию getUsersByCondition
	switch t := condition; t {
	case "age > 18":
		youngUsers = newUsers.Filter(func(user *User) bool {
			return user.Age > 18
		})
		return youngUsers.Items
	case "age < 18":
		youngUsers = newUsers.Filter(func(user *User) bool {
			return user.Age < 18
		})
		return youngUsers.Items
	case "age >= 18":
		youngUsers = newUsers.Filter(func(user *User) bool {
			return user.Age >= 18
		})
		return youngUsers.Items
	case "age <= 18":
		youngUsers = newUsers.Filter(func(user *User) bool {
			return user.Age <= 18
		})
		return youngUsers.Items
	case "age = 18":
		youngUsers = newUsers.Filter(func(user *User) bool {
			return user.Age == 18
		})
		return youngUsers.Items
	default:
		return nil
	}
}

func getUsersByAge(users []*User, age int) []*User {
	newUsers := NewUserCollection(users)
	var youngUsers *UserCollection
	youngUsers = newUsers.Filter(func(user *User) bool {
		return user.Age == age
	})
	return youngUsers.Items
}

func getUsersByNickName(users []*User, nickName string) []*User {
	newUsers := NewUserCollection(users)
	var youngUsers *UserCollection
	youngUsers = newUsers.Filter(func(user *User) bool {
		return user.Name == nickName
	})
	return youngUsers.Items
}

func getUsersUniqueNickName(users []*User) []*User {
	newUsers := NewUserCollection(users)
	var youngUsers *UserCollection
	youngUsers = newUsers.UniqByName()
	return youngUsers.Items
}
