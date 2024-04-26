package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var res string
	res = contextWithTimeout(context.Background(), 1*time.Second, 2*time.Second)
	fmt.Println(res)
	res = contextWithTimeout(context.Background(), 2*time.Second, 1*time.Second)
	fmt.Println(res)
}

func contextWithTimeout(ctx context.Context, contextTimeout time.Duration, timeAfter time.Duration) string {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()
	select {
	case <-ctxWithTimeout.Done():
		return "превышено время ожидания контекста"
	case <-time.After(timeAfter):
		return "превышено время ожидания"
	}
}
