package main

import (
	"testing"
	"web"
)

func TestUserNameAndPasswordTest(t *testing.T) {
	username := "tester"
	password := "1234"
	if !web.CheckLogin(username, password) {
		t.Error("fail")
	}

}
