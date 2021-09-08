package main

import (
	"testing"
)

func error(t *testing.T, got string, want string) {
	t.Errorf("got %q want %q", got, want)
}

func TestHello(t *testing.T) {
	got := Hello("Pablo")
	want := "Hello, Pablo"

	if got != want {
		error(t, got, want)
	}
}
