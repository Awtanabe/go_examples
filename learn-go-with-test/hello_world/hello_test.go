package main

import "testing"

// 基本
// func TestHello(t *testing.T) {
// 	// mainパッケージで同じなので、Hello関数呼べる
// 	// 実行
// 	got := Hello("Chris")
// 	// 期待
// 	want := "Hello, Chris"

// 	// 実行比較
// 	if got != want {
// 		// これはテストは止まらない。Fatalの場合だけ
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }


func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saylong hello to people", func (t *testing.T) {
		name := "Chris"
		got := Hello(name, "English")
		want := "Hello, " + name


		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello workd wehn an empty string", func (t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})


	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}