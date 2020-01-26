package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006"))
	fmt.Println("01-01-2020")
}
