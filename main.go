package main

import (
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

	go do(done)

	//!!! dont' doWithContext this (can make blocking operation) !!!
	//go do(done)
	//<-done

	//recommended (avoid blocking forever by handling timeout)
	select {
	case <-done:
	case <-time.After(time.Second):
		fmt.Println("timeout when receiving result")
	}

	fmt.Println("end app")
}

func do(callback chan error) {
	time.Sleep(time.Second)
	callback <- nil
}

/*func main() {
	fmt.Println("start app")

	time.Sleep(time.Second)

	done := make(chan error)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go doWithContext(ctx, done)
	fmt.Println(<-done)

	fmt.Println("end app")
}

//doWithContext function handle timeout by following the context
func doWithContext(ctx context.Context, callback chan error) {
	select {
	case <-ctx.Done():
		callback <- ctx.Err()
	case <-time.After(30 * time.Second):
		callback <- nil
		fmt.Println("done")
	}
}*/
