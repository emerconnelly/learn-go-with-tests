package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Emerson")
	want := "Hello, Emerson"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
