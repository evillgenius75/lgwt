package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	}

	t.Run("Hello with a specific name", func(t *testing.T) {
		name := "Eddie"
		got := Hello(name, "")
		want := "Hello, Eddie"

		assertMessage(t, got, want)

	})

	t.Run("Say Hello, World if empty param:", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Eduardo", "Spanish")
		want := "Hola, Eduardo"
		assertMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jesse", "French")
		want := "Bonjour, Jesse"
		assertMessage(t, got, want)
	})

}
