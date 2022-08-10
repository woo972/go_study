package main

import "fmt"

func main() {
	a := 1
	b := a

	fmt.Println(a, &a)
	fmt.Println(b, &b)

	b++

	fmt.Println(a)
	fmt.Println(b)

	a += 3

	fmt.Println(a)
	fmt.Println(b)

	c := 1
	d := &c

	fmt.Println(c, &c)
	fmt.Println(d, &d)

	//d++  // 참조 값에 대해서 재할당이나 연산이 불가능함

	c += 3

	fmt.Println(c)
	fmt.Println(d, &d, *d)

	*d++ // 역참조를 통해 값을 변경시키면 값이 모두 바뀌게 된다
	fmt.Println(c)
	fmt.Println(d, *d)

	add := func(x, y int) int {
		x++
		y++
		return x + y
	}

	a1 := 1
	b1 := 2
	fmt.Println(add(a1, b1), a1, b1)

	add2 := func(x int, y *int) int {
		x++
		*y++
		return x + *y
	}

	a2 := 1
	b2 := 2
	fmt.Println(add2(a2, &b2), a2, b2)
}
