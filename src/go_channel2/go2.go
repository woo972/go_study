package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- time.Now().String()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			ch2 <- "test"
			time.Sleep(2 * time.Second)
		}
	}()

	// for select case를 통해 각 채널읠 값을 수신 대기 할 수 있다
	for {
		select {
		case now := <-ch1:
			fmt.Println(now)
		case test := <-ch2:
			fmt.Println(test)

		}
	}
}
