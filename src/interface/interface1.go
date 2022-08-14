package main

import "fmt"

type Dog struct {
	animalType string
	age        int
}

// 인터페이스의 메서드를 구현한 리시버
func (animal Dog) bite() {
	fmt.Printf("%s is bite\n", animal.animalType)
}

type Cat struct {
	animalType string
	age        int
}

func (animal Cat) bite() {
	fmt.Printf("%s is bite\n", animal.animalType)
}

// 인터페이스에서 메서드 시그니처만 정의
type Behavior interface {
	bite()
}

func main() {
	dog := Dog{"dog", 10}
	dog.bite()

	var interface1 Behavior = Behavior(dog) // 인터페이스에 할당한 뒤 호출
	interface1.bite()

	cat := Cat{
		"cat",
		12,
	}
	cat.bite()

	interface2 := Behavior(cat)
	interface2.bite()

	interface3 := []Behavior{dog, cat}
	for i, v := range interface3 {
		interface3[i].bite()
		v.bite()
	}
}
