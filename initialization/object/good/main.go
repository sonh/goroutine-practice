package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// 1) Các goroutine làm tác vụ có liên quan đến nhau nên được quản lý và tập trung bởi một object
// 2) Object đó có nhiệm vụ quản lý lifecycle của các goroutines bên trong nó và cung cấp cho caller
//    phương thức để tương tác và quản lý lifcycle đó.

type GetDataResponse struct {
	data []int
	err  error
}

type CrawlerManager struct {
	data chan int

	getDataRequest chan chan []int

	stop chan struct{}
}

func NewCrawler() *CrawlerManager {
	crawler := &CrawlerManager{
		data:           make(chan int, 1000),
		getDataRequest: make(chan chan []int, 1000),
		stop:           make(chan struct{}),
	}

	//run goroutines bind with this object
	go crawler.run()

	return crawler
}

func (crawler *CrawlerManager) run() {
	var data []int

	go googleCrawler(crawler.data, crawler.stop)
	go facebookCrawler(crawler.data, crawler.stop)
	go facebookCrawler(crawler.data, crawler.stop)
	go facebookCrawler(crawler.data, crawler.stop)
	go facebookCrawler(crawler.data, crawler.stop)
	go facebookCrawler(crawler.data, crawler.stop)

	for {
		select {
		case crawledData := <-crawler.data:
			data = append(data, crawledData)
		case getDataCallback := <-crawler.getDataRequest:
			fmt.Println(data)
			getDataCallback <- data
		}
	}
}

func googleCrawler(crawledData chan int, stop chan struct{}) {
	for {
		time.Sleep(time.Second)
		data := rand.Intn(100000)

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

func facebookCrawler(crawledData chan int, stop chan struct{}) {
	for {
		time.Sleep(time.Second)
		data := rand.Intn(100000)

		select {
		case <-stop:
			return
		default:
			select {
			case crawledData <- data:
				fmt.Println("crawled data from B, saved")
				//......................
				//memory leak
			}
		}
	}
}

func (crawler *CrawlerManager) GetCrawledData(ctx context.Context) *GetDataResponse {
	getDataRequest := make(chan []int)
	crawler.getDataRequest <- getDataRequest

	select {
	case <-ctx.Done():
		return &GetDataResponse{
			err: ctx.Err(),
		}
	case crawledData := <-getDataRequest:
		fmt.Println(crawledData)
		return &GetDataResponse{
			data: crawledData,
		}
	}
}

func (crawler *CrawlerManager) Stop() {
	close(crawler.stop)
}

//endpoint
func main() {
	crawler := NewCrawler()

	time.Sleep(3 * time.Second)
	response := crawler.GetCrawledData(context.Background())

	if response.err != nil {
		log.Fatalln(response.err)
	}
	fmt.Println("data: ", response.data)

	crawler.Stop()
	fmt.Println("graceful shutdown crawler")
}
