package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, errors.New("need greather or equals 0") // errors.New(strings) 를 통해 에러를 생성할 수 있다
	}
	return math.Sqrt(input), nil
}

func main() {
	if rslt, err := Sqrt(-1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rslt)
	}
}
