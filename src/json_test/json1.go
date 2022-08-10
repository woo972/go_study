package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main() {
	user := struct {
		name string
		age  int
	}{"aaa", 10}

	fmt.Println(user, reflect.TypeOf(user))

	json.NewEncoder(os.Stdout).Encode(user)

}
