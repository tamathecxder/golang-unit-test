package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHelloAziziAssert(t *testing.T) {
	result := SayHello("Zee")

	assert.Equal(t, "Hello, Azizi", result, "Result must be 'Hello, Azizi'")

	fmt.Println("TestSayHelloAziziAssert DONE CUY")
}

func TestSayHelloShaniRequire(t *testing.T) {
	result := SayHello("Shani")

	require.Equal(t, "Hello, Shani", result, "Result must be 'Hello, Shani'")

	fmt.Println("TestSayHelloAziziAssert DONE CUY")
}

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
