package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known string", func(t *testing.T) {
		actual, _ := dictionary.Search("test")

		expected := "this is just a test"

		assertStrings(t, actual, expected)
	})

	t.Run("unknown string", func(t *testing.T) {
		_, err := dictionary.Search("demo")

		if err == nil {
			t.Fatal("Expect to return an error.")
		}

		assertError(t, err, ErrorKeyNotFound)
	})
}

func assertError(t testing.TB, actual, expected error) {
	t.Helper()
	if actual != expected {
		t.Errorf("Got %q but expect %q", actual, expected)
	}
}

func assertStrings(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Got %q but expect %q", actual, expected)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		key := "hello"
		value := "dictionary"

		dict.Add("hello", "dictionary")

		assertValue(t, dict, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "word"
		value := "existing value"
		dict := Dictionary{key: value}

		err := dict.Add(key, "new value")
		assertError(t, err, ErrorKeyExisting)
		assertValue(t, dict, key, value)
	})
}

func assertValue(t testing.TB, dict Dictionary, key, value string) {
	t.Helper()

	actual, err := dict.Search(key)

	if err != nil {
		t.Fatal("Key should exist: ", err)
	}

	assertStrings(t, actual, value)
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "hello"
		value := "not world"
		dict := Dictionary{key: value}

		expected := "world"
		dict.Update(key, expected)

		assertValue(t, dict, key, expected)
	})

	t.Run("not existing key", func(t *testing.T) {
		word := "hello"
		def := "world"

		dict := Dictionary{}

		err := dict.Update(word, def)

		assertError(t, err, ErrorWordNotExisting)
	})
}

func TestDelete(t *testing.T) {
	word := "key"

	dict := Dictionary{word: "definition"}

	dict.Delete(word)

	_, err := dict.Search(word)

	if err != ErrorKeyNotFound {
		t.Errorf("Expect %q to be deleted", word)
	}
}
