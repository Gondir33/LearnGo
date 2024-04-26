package main

import (
	"errors"
	"sync"
)

type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type CheckingAccount struct {
	balance float64
	mutex   sync.Mutex
}

func (c *CheckingAccount) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("negative amount sad :(")
	}
	c.mutex.Lock()
	c.balance += amount
	c.mutex.Unlock()
	return nil
}

func (c *CheckingAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return errors.New("negative amount sad :(")
	}
	if amount > c.balance {
		return errors.New("not enough money")
	}
	c.mutex.Lock()
	c.balance -= amount
	c.mutex.Unlock()
	return nil
}

func (c *CheckingAccount) Balance() float64 {
	return c.balance
}

type SavingsAccount struct {
	balance float64
	mutex   sync.Mutex
}

func (s *SavingsAccount) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("negative amount sad :(")
	}
	s.mutex.Lock()
	s.balance += amount
	s.mutex.Unlock()
	return nil
}

func (s *SavingsAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return errors.New("negative amount sad :(")
	}
	if s.balance-amount < 1000 {
		return errors.New("money will be less than 1000, can't do it")
	}
	s.mutex.Lock()
	s.balance -= amount
	s.mutex.Unlock()
	return nil
}

func (s *SavingsAccount) Balance() float64 {
	return s.balance
}
