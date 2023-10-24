package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Yudist", func(t *testing.T) {
		result := SayHello("Yudist")

		require.Equal(t, "Hello, Yudist", result, "Result must be 'Hello, Yudist'")
	})

	t.Run("Tama", func(t *testing.T) {
		result := SayHello("Tama")

		require.Equal(t, "Hello, Tama", result, "Result must be 'Hello, Tama'")
	})
}

func TestMain(t *testing.M) {
	// before
	fmt.Println("Before Unit Test")

	t.Run()

	// after
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run unit test in Windows Operating System")
	}

	result := SayHello("Asep")
	require.Equal(t, "Hello, Asep", result, "Result must be 'Hello, Asep'")
}

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
