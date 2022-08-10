package main

import "fmt"

func main() {
	// 조회 및 순회

	map1 := map[string]string{
		"daum":  "daum.net",
		"naver": "naver.com",
	}

	fmt.Println(map1)

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	for _, v := range map1 {
		fmt.Println(v)
	}

	map1["new"] = "old"
	map1["daum"] = "hanmail.net"
	fmt.Println(map1)

	value, ok := map1["new"] // 두번째 리턴값으로 키 존재 유무 확인
	fmt.Println(value, ok)

	delete(map1, "daum")

	if value, ok := map1["new"]; ok {
		fmt.Println(value)
	} else {
		fmt.Print("no value")
	}

}
