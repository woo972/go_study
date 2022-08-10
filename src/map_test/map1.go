package main

import "fmt"

func main() {
	// 맵은 레퍼런스 타입(참조값 전달)
	// 참조타입은 키로 사용이 불가능
	// make 함수 및 축약(리터럴)로 초기화가 가능

	/* 맵 선언 */
	var map1 map[string]int = make(map[string]int)
	var map2 = make(map[string]int)
	map3 := make(map[string]int)

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	map4 := map[string]int{} // Json 형태 <- Json a = {"a":1 , "b":2}
	map4["hello"] = 1
	map4["world"] = 2

	map5 := map[string]int{
		"hello": 3,
		"world": 4,
	}

	fmt.Println(map4)
	fmt.Println(map5)

	map6 := make(map[string]int, 1)
	map6["hello"] = 1
	map6["world"] = 2

	fmt.Println(map6)
}
