package main

import (
	"testing"
)

func TestUserNameAndPasswordTest(t *testing.T) {
	username := "tester"
	password := "1234"
	if !CheckLogin(username, password) {
		t.Error("fail")
	}

}
