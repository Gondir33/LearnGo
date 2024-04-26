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
	value map[string]interface{}
	mutex sync.Mutex
}

func NewCache() *Cache {
	return &Cache{value: make(map[string]interface{}), mutex: sync.Mutex{}}
}

func (c *Cache) Set(key string, user *User) {
	c.mutex.Lock()
	c.value[key] = user
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) interface{} {
	c.mutex.Lock()
	user := c.value[key]
	c.mutex.Unlock()
	return user
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func GetUser(i interface{}) *User {
	return i.(*User)
}

func main() {
	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprint("user-", i),
		})
	}

	time.Sleep(time.Second)

	for i := 0; i < 100; i++ {
		go func(i int) {
			raw := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			fmt.Println(GetUser(raw))
		}(i)
	}
}
