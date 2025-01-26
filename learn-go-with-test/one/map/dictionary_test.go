package main

import "testing"

func TestSearch(t *testing.T) {
	// dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		if got, err := dictionary.Search("test"); err != nil {
			t.Fatal("")
		} else {
			want := "this is just a test"
			assertStrings(t, got, want)		
		}
		
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("yahoo")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertStrings(t, err.Error(), want)	
	})
}

// ⭐️参照型がもたらす落とし穴は、マップがnil値になる可能性があることで
func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"

	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
			t.Fatal("should find added word:", err)
	}

	if definition != got {
			t.Errorf("got %q want %q", got, definition)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
			t.Errorf("got %q want %q", got, want)
	}
}