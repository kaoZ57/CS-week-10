package main

import (
	"fmt"
	"time"
)

func say(txt string) {
	for i := 0; i < 3; i++ {
		fmt.Println(time.Now(), ":", i, ":", txt)
		time.Sleep(time.Microsecond)
	}
}

func main() {
	go say("Hello")
	go say("say")
	var input string
	fmt.Scanln(&input)
}
