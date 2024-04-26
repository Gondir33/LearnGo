package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Product struct {
	Name      string
	Price     float64
	CreatedAt time.Time
	Count     int
}

func (p Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %f, Count: %v", p.Name, p.Price, p.Count)
}

func generateProducts(n int) []Product {
	gofakeit.Seed(time.Now().UnixNano())
	products := make([]Product, n)
	for i := range products {
		products[i] = Product{
			Name:      gofakeit.Word(),
			Price:     gofakeit.Price(1.0, 100.0),
			CreatedAt: gofakeit.Date(),
			Count:     gofakeit.Number(1, 100),
		}
	}
	return products
}

type ByCount []Product

func (b ByCount) Len() int           { return len(b) }
func (b ByCount) Less(i, j int) bool { return b[i].Count < b[j].Count }
func (b ByCount) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type ByPrice []Product

func (b ByPrice) Len() int           { return len(b) }
func (b ByPrice) Less(i, j int) bool { return b[i].Price < b[j].Price }
func (b ByPrice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type ByCreatedAt []Product

func (b ByCreatedAt) Len() int           { return len(b) }
func (b ByCreatedAt) Less(i, j int) bool { return b[i].CreatedAt.Unix() < b[j].CreatedAt.Unix() }
func (b ByCreatedAt) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

/*
func main() {
	products := generateProducts(10)

	fmt.Println("Исходный список:")
	fmt.Println(products)

	// Сортировка продуктов по цене
	sort.Sort(ByPrice(products))
	fmt.Println("\nОтсортировано по цене:")
	fmt.Println(products)

	// Сортировка продуктов по дате создания
	sort.Sort(ByCreatedAt(products))
	fmt.Println("\nОтсортировано по дате создания:")
	fmt.Println(products)

	// Сортировка продуктов по количеству
	sort.Sort(ByCount(products))
	fmt.Println("\nОтсортировано по количеству:")
	fmt.Println(products)
}
*/
