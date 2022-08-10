package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "string"
	var str2 string = "test"

	fmt.Println(str1 + str2)

	str3 := []string{} // string slice
	str3 = append(str3, str1)
	str3 = append(str3, str2)
	fmt.Println(str3)

	fmt.Println(strings.Join(str3, ", "))

}
