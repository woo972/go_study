package main

import (
	"fmt"
	"time"
)

type Test struct {
	count int
}

func (t *Test) inc(c chan int) {

	input := <-c
	t.count += input
	fmt.Printf("input: %v, result: %v\n", input, t.count)

}

func main() {

	t := Test{count: 0}
	ch1 := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			ch1 <- i
			t.inc(ch1)
			time.Sleep(5 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println()
			time.Sleep(10 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
}
