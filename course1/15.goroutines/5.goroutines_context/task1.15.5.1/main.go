package main

import (
	"context"
	"time"
)

func main() {
	var res string
	res = contextWithDeadline(context.Background(), 1*time.Second, 2*time.Second)
	println(res)
	res = contextWithDeadline(context.Background(), 2*time.Second, 1*time.Second)
	println(res)
	/* Output:
	context deadline exceeded
	time after exceeded
	*/
}

func contextWithDeadline(ctx context.Context, contextDeadline time.Duration, timeAfter time.Duration) string {
	// Создаем контекст с установленным временем истечения
	ctxWithDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(contextDeadline))
	defer cancel()

	select {
	case <-ctxWithDeadline.Done():
		// Если контекст истек, возвращаем соответствующую строку
		return "context deadline exceeded"
	case <-time.After(timeAfter):
		// Если время (timeAfter) истекло, возвращаем соответствующую строку
		return "time after exceeded"
	}
}
