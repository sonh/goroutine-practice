package main

import (
	"fmt"
	"time"
)

// bắt đầu lúc nào và kết thúc nào (life-cycle of goroutine)
// handle timeout cho các operation tương tác với channel

//Đối với function không return giá trị:
// 1) Với function không return giá trị, tổng thể của function luôn viết ở dạng synchronous (synchronous function)
// 2) Nếu code bên trong một synchronous function chạy async, cần phải có callback để caller xác định được khi nào function sẽ hoàn tất
// 3) Caller nên là người init callback và pass xuống các asynchronous function.

//Tại sao:
// 1) Bởi vì có thể dễ dàng execute một synchronous function ở dạng asynchronous với từ khóa "go" để execute
//    nó trên một goroutine mới, nhưng khó có thể execute một asynchronous function ở dạng synchronous mà ko
//    phải can thiệp vào logic của nó
// 2) Để caller không cần phải quá bận tâm về logic lẫn lifecycle bên trong function, chỉ cần tập trung vào input và output

/*func doSynchronously() {

	//Simulate the workload
	time.Sleep(time.Second)

	fmt.Println("done")
}

func main() {
	//GO
	fmt.Println("---------------")
	go doSynchronously()
	fmt.Println("---------------")
}*/

/*func doAsynchronously() {
	//Simulate the workload
	go func() {
		time.Sleep(time.Second)
		fmt.Println("done")
	}()
}

func main() {
	fmt.Println("---------------")
	go doAsynchronously()
	fmt.Println("---------------")
}*/

/*func doAsynchronously() chan struct{} {
	//use channel as callback
	done := make(chan struct{})

	//Simulate the workload
	go func() {
		time.Sleep(time.Second)
		fmt.Println("done")

		//notify callback
		done <- struct{}{}
	}()
	return done
}

func main() {
	fmt.Println("---------------")
	<-doAsynchronously()
	fmt.Println("---------------")
}*/

//Function không cần phải lo về việc init và return callback, tập trung vào logic của nó
func doAsynchronously(done chan struct{}) {
	//Simulate the workload
	go func() {
		time.Sleep(time.Second)
		fmt.Println("done")

		//notify callback
		done <- struct{}{}
	}()
}

func main() {
	fmt.Println("---------------")
	done := make(chan struct{})
	go doAsynchronously(done)
	<-done
	fmt.Println("---------------")
}
