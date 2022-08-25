package main

import (
	"context"
	"fmt"
	"time"
)

// 1) Luôn luôn xử lý timeout khi send/receive với channel

//channel is using to comminuate between goroutines

//G0 -> G1 -> G2 -> .... -> G(n)
//endpoint
func main() {
	fmt.Println("start app")

	time.Sleep(time.Second)

	done := make(chan error)

	ctx := context.Background()

	go do(ctx, done)

	//<-done // không khuyến khích (blocking operation)

	//
	select {
	case <-done:
	case <-time.After(time.Second):
		fmt.Println("timeout when receiving result")
	}

	fmt.Println("end app")
}

// timeout: 1 second
func do(ctx context.Context, callback chan error) {
	select {
	case <-ctx.Done():
		callback <- ctx.Err()
	case <-time.After(30 * time.Second):
		callback <- nil
		fmt.Println("done")
	}
}
