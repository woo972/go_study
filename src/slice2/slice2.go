package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
	  슬라이스 추가
	*/
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := []int{7, 8, 9}

	s1 = append(s1, 10, 11)
	fmt.Println(s1)

	s2 = append(s1, s2...) // 슬라이스를 추가할 때는 ...을 붙여주어야 한다
	fmt.Println(s2)

	s3 = append(s2, s3[1:3]...)
	fmt.Println(s3)

	s4 := make([]int, 5) // 길이, 용량이 5인 슬라이스 생성
	for i := 0; i < 20; i++ {
		// s4[i] = i // 이렇게 하면 out of bound 에러가 발생한다
		s4 = append(s4, i)
		fmt.Printf("len: %d cap: %d val: %v\n", len(s4), cap(s4), s4) // 메모리 용량을 넘어설 때마다 2배씩 증가시킨다. 최적화하려면 적절히 용량을 설정할 것
	}

	/*
	  슬라이스 추출
	*/
	fmt.Println(s3[:])   // 전체
	fmt.Println(s3[1:])  // index:1 부터
	fmt.Println(s3[:3])  // index: 3-1=2 까지
	fmt.Println(s3[1:3]) // index: 1 ~ 2 까지
	fmt.Println(s3[:len(s3)])

	/*
	  슬라이스 정렬
	*/
	s5 := []int{3, 5, 1, 2}
	fmt.Println("sorted? ", sort.IntsAreSorted(s5))
	sort.Ints(s5)
	fmt.Println("sorted ", s5)

	s6 := []string{"a", "c", "b"}
	fmt.Println("sorted? ", sort.StringsAreSorted(s6))
	sort.Strings(s6)
	fmt.Println("sorted ", s6)

	/*
	  슬라이스 복사
	*/
	s7 := []int{1, 2, 3}
	s8 := s7
	s8[0] = 10
	fmt.Println(s7, s8) // 슬라이스를 재할당하면 참조로 복사된다

	arr1 := [3]int{2, 2, 2}
	s9 := arr1[:]
	s9[0] = 10
	fmt.Println(arr1, s9) // 배열을 슬라이스로 재할당하면 참조로 복사된다

	s10 := []int{0, 1, 2}
	s11 := make([]int, 2) // 복사할 타겟에 메모리가 할당되어 있어야 하며, 그 만큼만 복사된다.
	copy(s11, s10)        // 타겟, 복사할 원본
	s11[0] = 10
	fmt.Println(s10, s11, len(s11), cap(s11)) // COPY를 이용하면 값만 복사할 수 있다.

}
