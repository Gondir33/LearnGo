package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID   int
	Name string
}

type Cache struct {
	cash  map[string]User
	mutex sync.Mutex
}

func NewCache() *Cache {
	return &Cache{cash: make(map[string]User), mutex: sync.Mutex{}}
}

func (c *Cache) Set(key string, user *User) {
	c.mutex.Lock()
	c.cash[key] = *user
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) *User {
	c.mutex.Lock()
	user := c.cash[key]
	c.mutex.Unlock()
	return &user
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func main() {
	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprint("user-", i),
		})
	}

	time.Sleep(1 * time.Second)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(cache.Get(keyBuilder("user", strconv.Itoa(i))))
		}(i)
	}
}
