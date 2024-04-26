package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	products := generateProducts(10)

	sort.Sort(ByPrice(products))
	assert.Equal(t, products, products)

	sort.Sort(ByCreatedAt(products))
	assert.Equal(t, products, products)

	sort.Sort(ByCount(products))
	assert.Equal(t, products, products)

	got := products[0].String()
	want := fmt.Sprintf("Name: %s, Price: %f, Count: %v", products[0].Name, products[0].Price, products[0].Count)
	assert.Equal(t, want, got)
}
