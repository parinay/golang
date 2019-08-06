package main

import "testing"

func TestHello(t *testing.T) {
	got := helloWorld()

	want := "Hello, World!"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
