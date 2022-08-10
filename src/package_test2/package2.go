package main

import (
	"fmt"
	"package_test2/library1"
)

var (
	i int32 = 11
)

func main() {
	fmt.Println(i, "는 10보다 큰 수인가?:", library1.CheckNum(11))
}
