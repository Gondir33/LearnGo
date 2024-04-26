package main

type Customer struct {
	ID      int
	Name    string
	Account Account
}

type CustomerOption func(*Customer)

func WithName(name string) CustomerOption {
	return func(c *Customer) {
		c.Name = name
	}
}

func WithAccount(account Account) CustomerOption {
	return func(c *Customer) {
		c.Account = account
	}
}

func NewCustomer(id int, options ...CustomerOption) *Customer {
	c := &Customer{ID: id}

	for _, option := range options {
		option(c)
	}

	return c
}

/*
func main() {
	savings := &SavingsAccount{}
	savings.Deposit(1000)

	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))

	err := customer.Account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Customer: %v, Balance: %v\n", customer.Name, customer.Account.Balance())
}
*/
