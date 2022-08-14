package main

import "fmt"

type Car struct { // 대문자 선언시 패키지 외부에서도 참조가 가능하다 (소문자는 패키지 내에서만 참조)
	name  string
	price int
}

// 일반 함수
func GetPrice(c Car, input int) int {
	return c.price + input
}

func GetPricePointer(c *Car, input int) int {
	return c.price + input
}

// 리시버1: 구조체와 함수를 바인딩한 형태
func (c Car) GetPrice2(input int) int {
	return c.price + input
}

// 리시버2: 포인터를 받아서 멤버변수 값을 바꿀수도 있다 -> 리시버로 선언한 경우 인자로 넘길 때 포인터형으로 정의하지 않아도 내부적으로 처리한다
func (c *Car) UpdatePrice(input int) {
	c.price = input
}

func add(i int) int {
	i++
	return i
}

type customInt int // 사용자 정의 타입

func varFunc(input ...int) int {
	var sum int
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum
}

func main() {

	kia := Car{name: "k3", price: 1000}
	audi := Car{"a7", 2000}

	fmt.Printf("%v, %p\n", kia, &kia)
	fmt.Printf("%v, %p\n", audi, &audi)

	fmt.Println(GetPrice(kia, 3000))
	fmt.Println(kia.GetPrice2(4000))

	kia.UpdatePrice(10)
	fmt.Println(kia)

	var a int = 1
	fmt.Println(add(a))
	var b customInt = 1
	// fmt.Println(add(b)) // 불가
	fmt.Println(add(int(b))) // 형변환을 해서 넣어야 한다

	fmt.Println(varFunc(1))

	arr := []int{1, 2}
	// fmt.Println(varFunc(arr)) // 안됨

	// 가변인자 슬라이스 형식으로 넣을 수 있다
	fmt.Println(varFunc(arr...))

	// 리시버, 일반 함수 포인터형 인자 비교
	normal := Car{"test1", 100}
	fmt.Println(GetPrice(normal, 1000)) // 값 전달임을 명시적으로 선언해야 한다

	normalPointer := Car{"test2", 1000}
	fmt.Println(GetPricePointer(&normalPointer, 1000)) // 포인터를 명시적으로 선언해야한다

	// 리시버로 선언시에는 포인터로 받을지 값으로 받을지 내부적으로 처리해준다
	reciever := Car{"test3", 10000}
	fmt.Println(reciever.GetPrice2(1000))

	reciever.UpdatePrice(10)
	fmt.Println(reciever.price)
}
