package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func handle1() {
	defer wg.Done()
	//for {
	// 	fmt.Println("handle1*************: request handler-1 before loop len-- ", len(requests))
	for req := range requests {
		//req := <-requests
		time.Sleep(100 * time.Millisecond)
		fmt.Println("handle1***************: request handling-1 num after processing -- ", req, len(requests))
	}
	//}

}

func handle() {
	defer wg.Done()
	//for {
	//	fmt.Println("handle ##################: request handler before loop len-- ", len(requests))

	for req := range requests {
		//req := <-requests
		time.Sleep(150 * time.Millisecond)
		fmt.Println("handle###############: request handling num after processing -- ", req, len(requests))
	}
	//}
}

func insertReq(start int, end int) {
	defer wg.Done()
	//for i := start; i <= end; i++ {
	//fmt.Println("trying to insert req -- ", start, len(requests))
	fmt.Println("Trying to process req -- ", start)
	limiter := time.Tick(wtime * time.Millisecond)
	<-limiter
	if xlimiter.Allow() == false {
		//fmt.Println("too many request per sec reached limit discarding req ^^^^^^^^^^^^^^^^^^ ", start, len(requests))
		fmt.Println("Too many request per sec, reached limit discarding failed req ^^^^^^^^^^^^^^^^^^ ", start)
		//wtime = 100
		//go handle1()
		//continue
		return
	}
	fmt.Println("Accepted req successfully -- ", start)

	//requests <- start

	//}
	//fmt.Println("total requests in the channel--- ", len(requests))
	return

}

var requests chan int
var wg sync.WaitGroup
var xlimiter = rate.NewLimiter(4, 10)
var wtime time.Duration

func main() {
	wtime = 50
	requests = make(chan int, 100)
	workers := 50
	// wg.Add(1)
	// go handle()
	// wg.Add(1)
	// go handle1()

	// Increment waitgroup counter and create go routines
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		//limiter := time.Tick(wtime * time.Millisecond)
		//<-limiter

		go insertReq(i, i)
		//wg.Add(1)
		//go insertReq(11, 20)

	}
	wg.Wait()
	fmt.Println("----------closing requests----------------- ")
	close(requests)

	fmt.Println("----------finished processing----------------- ")

	/*
	       burstyLimiter := make(chan time.Time, 3)

	       for i := 0; i < 3; i++ {
	           burstyLimiter <- time.Now()
	       }

	       go func() {
	           for t := range time.Tick(200 * time.Millisecond) {
	               burstyLimiter <- t
	           }
	       }()

	       burstyRequests := make(chan int, 5)
	       for i := 1; i <= 5; i++ {
	           burstyRequests <- i
	       }
	       close(burstyRequests)
	       for req := range burstyRequests {
	           <-burstyLimiter
	           fmt.Println("request", req, time.Now())
	   	}
	*/
}
