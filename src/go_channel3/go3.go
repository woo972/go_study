package main

import (
	"fmt"
	"time"
)

// 채널을 함수의 반환값으로 사용할 수 있다
func sumFrom1(num int) <-chan int {
	ch := make(chan int)
	sum := 0
	// 클로져를 이용해 채널에 값을 담는다
	go func() {
		for i := 1; i <= num; i++ {
			sum += i
		}
		ch <- sum
	}()
	return ch
}

// 채널을 함수의 인자로 받을 수 있다
func sumMinus1(sumFrom1Rslt <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		rslt := <-sumFrom1Rslt
		ch <- rslt - 1
	}()
	return ch
}

func main() {
	a := sumFrom1(10)
	b := sumFrom1(11)
	fmt.Println(<-sumMinus1(a))
	fmt.Println(<-sumMinus1(b))

	// a <- 1
	// b <- 1  // 수신 전용으로 설정되어 할당은 되지 않음

	time.Sleep(1 * time.Second)
}
