package main

import "fmt"

func main() {
	// 슬라이스는 길이가 가변이고, 레퍼런스 타입
	// 크기를 동적으로 할당 가능하다
	// 리스트 처럼 배열 보다는 이쪽을 많이 사용할 듯 함

	//생성방법 1
	var slice1 []int // nil
	slice2 := []int{}
	slice3 := []int{1, 2, 3}
	slice4 := [][]int{
		{1, 3, 5},
		{2, 4},
	}
	slice3[1] = 1000

	if slice1 == nil {
		fmt.Println("slice1 is nil")
	}

	fmt.Printf("1 %-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("2 %-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("3 %-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("4 %-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	// 생성방법 2
	var s1 []int = make([]int, 5, 10) // 길이는 5이고 용량은 10 (메모리 용량에 근접하게 실제 데이터를 사용하는 것이 좋다, 5개가 0으로 초기화)
	var s2 []int = make([]int, 5, 5)  // 길이와 용량을 5로 제한
	var s3 = make([]int, 5)
	s4 := make([]int, 5, 10)
	s5 := make([]int, 5)

	s1[2] = 1000

	fmt.Printf("%-5T %d %d %v\n", s1, len(s1), cap(s1), s1)
	fmt.Printf("%-5T %d %d %v\n", s2, len(s2), cap(s2), s2)
	fmt.Printf("%-5T %d %d %v\n", s3, len(s3), cap(s3), s3)
	fmt.Printf("%-5T %d %d %v\n", s4, len(s4), cap(s4), s4)
	fmt.Printf("%-5T %d %d %v\n", s5, len(s5), cap(s5), s5)

	// 배열은 값을 복사하는데 반해, 슬라이스는 메모리 주소 복사
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	arr1[0] = 1000
	fmt.Println(arr2[0])

	sl1 := []int{1, 2, 3}
	sl2 := sl1
	sl1[0] = 100
	fmt.Println(sl2[0])

	// 슬라이스 예외
	//sl3 := make([]int, 50, 100)
	//fmt.Println(sl3[50]) // out of bound

	// 슬라이스 순회
	for i, v := range sl1 {
		fmt.Println(i, v)
	}
}
