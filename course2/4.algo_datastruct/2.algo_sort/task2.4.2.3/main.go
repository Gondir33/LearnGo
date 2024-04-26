package main

type User struct {
	ID   int
	Name string
	Age  int
}

// Функция слияния двух отсортированных массивов пользователей
func Merge(arr1 []User, arr2 []User) []User {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	var i, j int
	res := make([]User, 0, len(arr1)+len(arr2))
	for i < len(arr1) && j < len(arr2) {
		if arr1[i].ID < arr2[j].ID {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	if i == len(arr1) && j != len(arr2) {
		for ; j < len(arr2); j++ {
			res = append(res, arr2[j])
		}
	}
	if i != len(arr1) && j == len(arr2) {
		for ; i < len(arr1); i++ {
			res = append(res, arr1[i])
		}
	}

	return res
}
