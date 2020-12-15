package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequestTime(ctx context.Context) {
	go WriteRedisTime(ctx)
	go WriteDatabaseTime(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}
func WriteRedisTime(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}
func WriteDatabaseTime(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 7*time.Second)
	go HandelRequestTime(ctx)
	time.Sleep(10 * time.Second)
}
