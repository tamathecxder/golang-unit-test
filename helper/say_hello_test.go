package helper

import (
	"fmt"
	"testing"
)

func TestSayHelloJihyo(t *testing.T) {
	result := SayHello("Jihyo")

	if result != "Hello, Jihyo" {
		t.Error("It must be 'Hello, Jihyo'")
	}

	fmt.Println("TestSayHelloJihyo DONE CUY")
}

func TestSayHelloAshel(t *testing.T) {
	result := SayHello("Ashel")

	if result != "Hello, Ashel" {
		t.Fatal("It must be 'Hello, Ashel'")
	}

	fmt.Println("TestSayHelloAshel DONE CUY")
}
