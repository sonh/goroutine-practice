package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

}

func googleCrawler(crawledData chan int, stop chan struct{}) {
	for {
		time.Sleep(time.Second)
		data := rand.Intn(100000)

		select {
		case <-stop:
			close(crawledData)
			return
		default:
			select {
			case crawledData <- data:
				fmt.Println("crawled data from B, saved")
			case <-time.After(time.Second):
				fmt.Println("timeout when saving crawled data")
			}
		}
	}
}
