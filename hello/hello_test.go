package main

import (
	"testing"
)

/*  rules for writing test files
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test - Test*
- The test function takes one argument only t *testing.T */

func TestHello(t *testing.T) {
	assertCorrectMessage := func(want string, got string, t *testing.T) {
		t.Helper() // ref: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#hello-world-again
		if want != got {
			t.Errorf("Wanted %q but instead got %q", want, got)
		}
	}

	t.Run("Hello, {NAME}", func(t *testing.T) {
		want := "Hello, Zubin!"
		got := Hello("Zubin", ENG)
		assertCorrectMessage(want, got, t)
	})

	t.Run("Hello, World -> no name passed", func(t *testing.T) {
		want := "Hello, World!"
		got := Hello("", ENG)

		assertCorrectMessage(want, got, t)
	})

	t.Run("Accepts Spanish in Hello()", func(t *testing.T) {
		got := Hello("Elodie", SPANISH)
		want := "Hola, Elodie!"

		assertCorrectMessage(want, got, t)
	})

	t.Run("Accepts French in Hello()", func(t *testing.T) {
		got := Hello("Rowena", FRENCH)
		want := "Bonjour, Rowena!"

		assertCorrectMessage(want, got, t)
	})

}
