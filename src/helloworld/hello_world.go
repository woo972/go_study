package main

import "fmt"

func main() {
	fmt.Println("<<hello world")

	fmt.Println("<<loop:")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("<<if")
	var str1 string
	str1 = "test"
	if str1 == "test" {
		fmt.Println(str1)
	}

	fmt.Println("<<var / const")
	const (
		c1 float32 = 0.1
		c2 string  = "abc"
	)
	// cant not assign
	// c1 = 0.2

	var (
		v1 float32 = 0.1
		v2 string
	)

	v1 = 0.2
	v2 = "abc"

	fmt.Println(v1, v2)

	fmt.Println("iota")
	const (
		i1 = iota
		i2
		i3
	)

	fmt.Println(i1, i2, i3)
}
