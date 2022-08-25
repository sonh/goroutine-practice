package main

import (
	"fmt"
	"math/rand"
	"time"
)

//endpoint
func main() {
	crawledData := make(chan int, 1000)
	stop := make(chan struct{})

	go googleCrawler(crawledData, stop)
	go facebookCrawler(crawledData, stop)
	go facebookCrawler(crawledData, stop)
	go facebookCrawler(crawledData, stop)
	go facebookCrawler(crawledData, stop)

	time.Sleep(3 * time.Second)
	close(stop)

	var data []int
	for d := range crawledData {
		data = append(data, d)
	}

	fmt.Println("data: ", data)
}

//repo
func googleCrawler(crawledData chan int, stop chan struct{}) {
	for {
		//simulate
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

func facebookCrawler(crawledData chan int, stop chan struct{}) {
	time.Sleep(time.Second)
	data := rand.Intn(100000)

	for {
		select {
		case <-stop:
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
