package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldJohn(t *testing.T) {
	result := HelloWorld("John")

	if result != "Hello John" {
		t.Fail()
	}
	fmt.Println("TestHelloWorldJohn Done")
}

func TestHelloWorldDow(t *testing.T) {
	result := HelloWorld("Dow")

	if result != "Hello Dow" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorldDow Done")
}
