package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	name   string
	price  int
	detail spec // 구조체를 중첩시킬 수 있다
}

type spec struct {
	height int
	widht  int
}

func (c *Car) getName() string {
	return c.name
}

func main() {

	// 익명 구조체를 만들 수 있다
	car1 := struct {
		name  string "[name of car]" // 각 필드에 태그를 달 수 있다 (주석 대신 사용 가능)
		price int    "[price of car]"
	}{"audi", 1000}
	fmt.Printf("%v / %#v\n", car1, car1) // %#v로 타입과 값을 모두 출력할 수 있다

	carReflect := reflect.TypeOf(car1)
	for i := 0; i < carReflect.NumField(); i++ {
		fmt.Printf("%s %s %s\n", carReflect.Field(i).Name, carReflect.Field(i).Type, carReflect.Field(i).Tag)
	}

	c1 := new(Car)
	c1.name = "a1"
	fmt.Println(c1.getName())

	c2 := Car{name: "a2"}
	fmt.Println(c2.getName())

	var c3 *Car = new(Car)
	c3.name = "a3"
	fmt.Println(c3.getName())
}
