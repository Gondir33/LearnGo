package main

import "sort"

type UserCollection struct {
	Items []*User
}

type SearchCallbackUser func(item *User) bool

//grizzly:replaceName New{{.Name}}Collection
func NewUserCollection(items []*User) *UserCollection {
	var collection UserCollection

	collection.Items = items

	return &collection
}

//grizzly:replaceName NewEmpty{{.Name}}Collection
func NewEmptyUserCollection() *UserCollection {
	return &UserCollection{}
}

func (c *UserCollection) Find(callback SearchCallbackUser) *User {
	for _, v := range c.Items {
		if callback(v) == true {
			return v
		}
	}

	return nil
}

func (c *UserCollection) Filter(callback SearchCallbackUser) *UserCollection {
	var newItems []*User

	for _, v := range c.Items {
		if callback(v) == true {
			newItems = append(newItems, v)
		}
	}

	return &UserCollection{Items: newItems}
}

func (c *UserCollection) MapToInt(callback func(item *User) int) []int {
	items := []int{}

	for _, v := range c.Items {
		items = append(items, callback(v))
	}

	return items
}

func (c *UserCollection) MapToString(callback func(item *User) string) []string {
	items := []string{}

	for _, v := range c.Items {
		items = append(items, callback(v))
	}

	return items
}

func (c *UserCollection) Push(item *User) *UserCollection {
	newItems := append(c.Items, item)

	return &UserCollection{Items: newItems}
}

func (c *UserCollection) Shift() *User {
	item := c.Items[0]
	c.Items = c.Items[1:]

	return item
}

func (c *UserCollection) Pop() *User {
	item := c.Items[len(c.Items)-1]
	c.Items = c.Items[:len(c.Items)-1]

	return item
}

func (c *UserCollection) Unshift(item *User) *UserCollection {
	newItems := append([]*User{item}, c.Items...)

	return &UserCollection{Items: newItems}
}

func (c *UserCollection) Len() int {
	return len(c.Items)
}

func (c *UserCollection) Get(index int) (model *User) {
	if index >= 0 && len(c.Items) > index {
		return c.Items[index]
	}

	return model
}

func (c *UserCollection) UniqByName() *UserCollection {
	collection := &UserCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *User) bool {
			return model.Name == item.Name
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *UserCollection) UniqByAge() *UserCollection {
	collection := &UserCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *User) bool {
			return model.Age == item.Age
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

func (c *UserCollection) UniqById() *UserCollection {
	collection := &UserCollection{}

	for _, item := range c.Items {
		searchItem := collection.Find(func(model *User) bool {
			return model.Id == item.Id
		})

		if searchItem == nil {
			collection = collection.Push(item)
		}
	}

	return collection
}

type byNameAsc []*User

func (a byNameAsc) Len() int           { return len(a) }
func (a byNameAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byNameAsc) Less(i, j int) bool { return a[i].Name < a[j].Name }

type byNameDesc []*User

func (a byNameDesc) Len() int           { return len(a) }
func (a byNameDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byNameDesc) Less(i, j int) bool { return a[i].Name > a[j].Name }

func (c *UserCollection) SortByName(mode string) *UserCollection {
	collection := &UserCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byNameDesc(collection.Items))
	} else {
		sort.Sort(byNameAsc(collection.Items))
	}

	return collection
}

type byAgeAsc []*User

func (a byAgeAsc) Len() int           { return len(a) }
func (a byAgeAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAgeAsc) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byAgeDesc []*User

func (a byAgeDesc) Len() int           { return len(a) }
func (a byAgeDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAgeDesc) Less(i, j int) bool { return a[i].Age > a[j].Age }

func (c *UserCollection) SortByAge(mode string) *UserCollection {
	collection := &UserCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byAgeDesc(collection.Items))
	} else {
		sort.Sort(byAgeAsc(collection.Items))
	}

	return collection
}

type byIdAsc []*User

func (a byIdAsc) Len() int           { return len(a) }
func (a byIdAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIdAsc) Less(i, j int) bool { return a[i].Id < a[j].Id }

type byIdDesc []*User

func (a byIdDesc) Len() int           { return len(a) }
func (a byIdDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIdDesc) Less(i, j int) bool { return a[i].Id > a[j].Id }

func (c *UserCollection) SortById(mode string) *UserCollection {
	collection := &UserCollection{Items: c.Items}

	if mode == "desc" {
		sort.Sort(byIdDesc(collection.Items))
	} else {
		sort.Sort(byIdAsc(collection.Items))
	}

	return collection
}

func (c *UserCollection) ForEach(callback func(item *User)) {
	for _, i := range c.Items {
		callback(i)
	}
}
