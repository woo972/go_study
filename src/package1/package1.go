package main

import (
	"fmt"
	"os"
)

func main() {
	// 패키지 이름 = 디렉토리 이름
	// 같은 패키지 내 = 소스파일들은 디렉토리명을 패키지명으로 사용
	// 네이밍 규칙 = 소문자 private, 대문자 public
	// main 함수가 들어 있는 파일은 프로그램 시작점

	var name string

	fmt.Println("name:")

	// 사용자 입력을 받아 변수에 저장한다
	fmt.Scanf("%s", &name)

	// COMMAND에 출력
	fmt.Fprintf(os.Stdout, "Hi %s\n", name)

}
