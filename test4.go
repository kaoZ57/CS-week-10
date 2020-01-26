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
