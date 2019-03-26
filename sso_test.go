package main

import (
	"fmt"
	"testing"
)

func TestInsert(test *testing.T) {
	_ = Register(
		"testEmail@ieee.org",
		"testPassword",
		"testData")
}

func TestLogin(test *testing.T) {
	fmt.Println(GetToken("testEmail@ieee.org",
		"testPassword"))
}
