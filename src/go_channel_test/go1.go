package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2) // string을 주고 받는 chan타입으로 channel 생성, 두번째 인자로 버퍼의 크기를 설정할 수 있다
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("channel input %dth:\n", i)
			ch <- "data" // 채널로 데이터를 보낸다
		}

		close(ch) // channel을 닫아주어야 한다.
	}()

	for v := range ch { // range channel을 통해 채널 데이터를 계속 받아올 수 있다
		fmt.Println(v)
	}

	value, ok := <-ch      // 두개의 값을 받을 수 있다. (실제 전송된 값, 상태)
	fmt.Println(value, ok) // 채널이 닫힌 뒤의 상태는 false다
}
