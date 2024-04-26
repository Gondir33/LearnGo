package main

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	res := make([]User, 0, cap(users))
	uniqueUsers := make(map[string]bool)
	for i := 0; i < len(users); i++ {
		_, ok := uniqueUsers[users[i].Nickname]
		if ok == false {
			uniqueUsers[users[i].Nickname] = true
			res = append(res, users[i])
		}
	}
	return res
}
