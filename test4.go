package main

import (
	"fmt"
	"sync"
	"time"
)

func sayjj(txt string, sleep time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(txt)
	time.Sleep(time.Microsecond * sleep)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go sayjj("hello", 2, &wg)
	go sayjj("Hi", 1, &wg)
	wg.Wait()
	fmt.Println("GoodBye")
}
