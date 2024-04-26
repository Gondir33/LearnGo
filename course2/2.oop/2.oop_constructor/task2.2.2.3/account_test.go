package main

import (
	"errors"
	"testing"
)

type TestAccountData struct {
	acc   Account
	input float64
	want  float64
	err   error
}

func TestDeposit(t *testing.T) {
	datas := []TestAccountData{{
		acc:   &SavingsAccount{},
		input: 1000,
		want:  1000,
	}, {
		acc:   &CheckingAccount{},
		input: 1000,
		want:  1000,
	}, {
		acc:   &CheckingAccount{},
		input: -1000,
		want:  0,
		err:   errors.New("negative amount sad :("),
	}, {
		acc:   &SavingsAccount{},
		input: -1000,
		want:  0,
		err:   errors.New("negative amount sad :("),
	},
	}
	for _, data := range datas {
		err := data.acc.Deposit(data.input)
		if err != nil {
			if err.Error() != data.err.Error() {
				t.Errorf("get %v, want %v", err, data.err)
			}
		}
		if data.acc.Balance() != data.want {
			t.Errorf("get %v; want %v", data.acc.Balance(), data.want)
		}
	}
}
func TestWithdraw(t *testing.T) {
	datas := []TestAccountData{{
		acc:   &SavingsAccount{balance: 1200},
		input: 100,
		want:  1100,
	}, {
		acc:   &CheckingAccount{balance: 1200},
		input: 100,
		want:  1100,
	}, {
		acc:   &SavingsAccount{balance: 1200},
		input: -200,
		want:  1200,
		err:   errors.New("negative amount sad :("),
	}, {
		acc:   &CheckingAccount{balance: 1200},
		input: -200,
		want:  1200,
		err:   errors.New("negative amount sad :("),
	}, {
		acc:   &SavingsAccount{balance: 1200},
		input: 800,
		want:  1200,
		err:   errors.New("money will be less than 1000, can't do it"),
	}, {
		acc:   &CheckingAccount{balance: 1200},
		input: 2000,
		want:  1200,
		err:   errors.New("not enough money"),
	},
	}
	for _, data := range datas {
		err := data.acc.Withdraw(data.input)
		if err != nil {
			if err.Error() != data.err.Error() {
				t.Errorf("get %v, want %v", err, data.err)
			}
		}
		if data.acc.Balance() != data.want {
			t.Errorf("get %v; want %v", data.acc.Balance(), data.want)
		}
	}
}

func TestBalance(t *testing.T) {
	datas := []TestAccountData{{
		acc:  &SavingsAccount{balance: 500},
		want: 500,
	}, {
		acc:  &CheckingAccount{balance: 500},
		want: 500,
	}}
	for _, data := range datas {
		data.acc.Withdraw(data.input)
		if data.acc.Balance() != data.want {
			t.Errorf("get %v; want %v", data.acc.Balance(), data.want)
		}
	}
}
