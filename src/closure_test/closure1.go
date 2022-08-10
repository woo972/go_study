package main

import "fmt"

func main() {
	// 참조하는 변수를 라이프사이클이 종료되도 삭제하지 않고 계속 참조하게 만든다
	// 오용할 경우 메모리 이슈가 생길 가능성

	// 지역변수는 함수가 종료되면 함께 사라진다
	isTargetToDisplay := func() bool {
		var isShow bool
		fmt.Println("init:", isShow)
		isShow = !isShow
		return isShow
	}
	fmt.Println(isTargetToDisplay())
	fmt.Println(isTargetToDisplay())

	// 전역변수는 함수가 종료되어도 유지된다
	// 그러나 다른 함수에서 값을 바꾸거나 기타 등등 문제가 생길 수 있다
	var isShow1 bool
	isTargetToDisplay2 := func() bool {
		fmt.Println("init:", isShow1)
		isShow1 = !isShow1
		return isShow1
	}
	fmt.Println(isTargetToDisplay2())
	fmt.Println(isTargetToDisplay2())

	// 즉시 실행 함수는 한 번 실행 뒤에 사라진다 (=> 초기화가 한번만 된다)
	// 그러나 변수를 클로져로 참조하기 때문에 변수는 메모리에 남는다
	isTargetToDisplay3 := func() func() bool {
		var isShow bool
		fmt.Println("init:", isShow)
		return func() bool {
			isShow = !isShow
			return isShow
		}
	}()
	fmt.Println(isTargetToDisplay3()) //
	fmt.Println(isTargetToDisplay3())
}
