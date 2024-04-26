package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Cacher interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) Cacher {
	return &cache{
		client: client,
	}
}

func (c *cache) Set(key string, value interface{}) error {

	err := c.client.Set(key, value, 5*time.Minute).Err()
	return err
}
func (c *cache) Get(key string) (interface{}, error) {
	var arg interface{}

	stringCmd := c.client.Get(key)
	err := stringCmd.Err()

	if err == redis.Nil {
		return 0, errors.New(fmt.Sprintf("not found by key %v", key))
	} else if err != nil {
		return 0, err
	}
	stringCmd.Scan(&arg)
	return arg, nil
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func main() {
	// Создание клиента Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	cache := NewCache(client)

	// Установка значения по ключу
	err := cache.Set("some:key", "value")
	if err != nil {
		panic(err)
	}

	// Получение значения по ключу
	value, err := cache.Get("some:key")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)

	user := &User{
		ID:   1,
		Name: "John",
		Age:  30,
	}
	// Установка значения по ключу
	err = cache.Set(fmt.Sprintf("user:%v", user.ID), user)
	if err != nil {
		panic(err)
	}

	// Получение значения по ключу
	value, err = cache.Get("key")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}
