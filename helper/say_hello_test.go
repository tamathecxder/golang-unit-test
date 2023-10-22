package helper

import "testing"

func TestSayHelloJihyo(t *testing.T) {
	result := SayHello("Jihyo")

	if result != "Hello, Jihyo" {
		panic("Result is not `Hello, Jihyo`")
	}
}

func TestSayHelloAshel(t *testing.T) {
	result := SayHello("Anin")

	if result != "Hello, Ashel" {
		panic("Result is not `Hello, Ashel`")
	}
}
