package main

import (
	"fmt"
	"time"
)

type User struct {
	ID      int
	Name    string
	Bonuses int
}

type BonusOperation struct {
	UserID int
	Amount int
}

var users = []*User{
	{ID: 1, Name: "Bob", Bonuses: 100},
	{ID: 2, Name: "Alice", Bonuses: 200},
	{ID: 3, Name: "Kate", Bonuses: 300},
	{ID: 4, Name: "Tom", Bonuses: 400},
	{ID: 5, Name: "John", Bonuses: 500},
}

func main() {
	bonusOperations := []BonusOperation{
		{UserID: 1, Amount: 10},
		{UserID: 2, Amount: 20},
		{UserID: 3, Amount: 30},
		{UserID: 4, Amount: 40},
		{UserID: 5, Amount: 50},
	}

	DeductBonuses(users, bonusOperations)
	for _, user := range users {
		// fmt.Printf
		fmt.Printf("User %s has %d bonuses\n", user.Name, user.Bonuses)
	}

}

// DeductBonuses - вычитает бонусы у пользователя
func DeductBonuses(users []*User, bonusesOperations []BonusOperation) {
	for i, _ := range users {
		go func(i int, user *User, bonusesOperation BonusOperation) { // исправить анонимную функцию, не удаляя ее
			// здесь мы как будто обращаемся во внешний сервис для списания бонусов
			user.Bonuses -= bonusesOperation.Amount
		}(i, users[i], bonusesOperations[i])
		time.Sleep(time.Duration(time.Millisecond))
	}

}
