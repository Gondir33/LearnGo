package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://httpbin.org/get"
	parallelRequest := 5
	requestCount := 50
	result := benchRequest(url, parallelRequest, requestCount)

	for i := 0; i < requestCount; i++ {
		statusCode := <-result
		if statusCode != 200 {
			panic(fmt.Sprintf("Ошибка при отправке запроса: %d", statusCode))
		}
	}

	fmt.Println("Все горутины завершили работу")
}

func benchRequest(url string, parallelRequest int, requestCount int) <-chan int {
	limit := make(chan int, parallelRequest)
	go func() {
		for i := 0; i < requestCount; i++ {
			statusCode, err := httpRequest(url)
			if err != nil {
				panic("err in requuest")
			}
			limit <- statusCode
		}
	}()
	return limit
}

func httpRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
