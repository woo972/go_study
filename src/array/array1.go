package main

import "fmt"

func main() {
	// cap() 배열, 슬라이스의 용량
	// len() 배열, 슬라이스의 개수

	var arr1 [3]int                   // 선언만하면 0으로 초기화 된다
	var arr2 [3]int = [3]int{1, 2, 3} // 크기를 지정하고 뒤에 이어서 중괄호를 사용해서 초기화
	var arr3 = [3]int{4, 5, 6}
	arr4 := [3]int{1, 1, 1}
	arr5 := [...]int{11, 1, 1, 1, 1, 1, 1} // 가변배열
	arr6 := [2][2]int{
		{1, 2},
		{3, 4}, // 쉼표를 적고 끝내야 한다
	}

	fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6)
	fmt.Printf("%-5T %d %v\n", arr6, len(arr6), arr6)

	arry1 := [...]string{"a", "b", "c"}
	for i := 0; i < len(arry1); i++ {
		fmt.Println(arry1[i])
	}

	// range를 이용한 루프
	for i, v := range arry1 { // i=index, v=value
		fmt.Println(i, "/", v)
	}

	for _, v := range arry1 {
		fmt.Println(v)
	}

	test1 := [2]int{1, 10}
	test2 := test1
	fmt.Println(test1, test2, &test1, &test2) // 이렇게 하면 포인터 주소값이 출력되지 않는다
	fmt.Printf("%p %v\n", &test1, test1)      // p=pointer, v=value
	fmt.Printf("%p %v\n", &test2, test2)      // 서로 다른 주소에 저장되어 값 복사가 일어나는 것을 알 수 있다
}
