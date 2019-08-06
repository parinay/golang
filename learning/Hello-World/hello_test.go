package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	// subtest-1
	t.Run("saying hello to people", func(t *testing.T) {
		got := helloWorld("Chris!")
		want := "Hello, Chris!"

		assertMessage(t, got, want)
	})

	// subtest-2
	t.Run("say hello world when an emptry string is supplied", func(t *testing.T) {
		got := helloWorld("")
		want := "Hello, World!"

		assertMessage(t, got, want)
	})
}
