package main

import (
	"fmt"
	"sync"
	"time"
)

func testSleepOne() {
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("DONE SLEEPING")
}

func testSleepTwo(waitTime int, resp chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Defer - Do this right before the function returns
	time.Sleep(time.Millisecond * time.Duration(waitTime))
	resp <- "DONE SLEEPING 2"
}

func main() {
	start := time.Now()
	testSleepOne()
	testSleepOne()
	fmt.Println("Time taken for sync code: ", time.Since(start))

	respChan := make(chan string)
	var wg sync.WaitGroup
	start = time.Now()
	wg.Add(2)
	go testSleepTwo(1000, respChan, &wg)
	go testSleepTwo(1500, respChan, &wg)
	go func() {
		wg.Wait()
		close(respChan)
	}()
	for resp := range respChan {
		fmt.Println(resp)
	}
	fmt.Println("Time taken for async code 1: ", time.Since(start))

	respChan = make(chan string)
	wg.Add(100)
	start = time.Now()
	for i := 0; i < 100; i++ {
		go testSleepTwo(1000+(i*10), respChan, &wg)
	}
	go func() {
		wg.Wait()
		close(respChan)
	}()
	for resp := range respChan {
		fmt.Println(resp)
	}
	fmt.Println("Time taken for async code 2: ", time.Since(start))
}
