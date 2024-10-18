package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	// key type can only be a comparable type
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		key := "test"
		got, _ := dictionary.Search(key)
		expected := "this is just a test"

		assertStrings(t, got, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		key := "unknown"
		_, err := dictionary.Search(key)

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrorNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q want %q given", got, expected)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		definition := "this is just a test"
		dictionary.Add(key, definition)

		assertDefinition(t, dictionary, key, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "henry"
		definition := "an engineer always work overtime"

		dictionary := Dictionary{key: definition}

		err := dictionary.Add(key, "never has workout experience")

		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, key, definition)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "henry"
		definition := "an engineer always work overtime"

		dictionary := Dictionary{word: definition}
		newDefinition := "an engineer always workouts 5 times/week"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "henry"
		definition := "an engineer always work overtime"

		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrorWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	word := "henry"
	definition := "an engineer always work overtime"

	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)

	_, err := dictionary.Search("henry")
	assertError(t, err, ErrorNotFound)
}
