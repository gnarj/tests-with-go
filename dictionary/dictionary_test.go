package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("testing")
		want := ErrNotFound

		if err == nil {
			t.Fatal("did not receive error when we need one")
		}

		assertError(t, err, want)
	})

}

func TestAdd(t *testing.T) {
	t.Run("add an item to a dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("try to add an item that already exists", func(t *testing.T) {
		definition := "yes"
		word := "exists"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update a key in a dictionary with a new value", func(t *testing.T) {
		key := "color"
		color := "white"
		dictionary := Dictionary{key: color}
		newColor := "blue"

		err := dictionary.Update(key, newColor)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newColor)
	})
	t.Run("update a key that doesn't exist", func(t *testing.T) {
		key := "color"
		color := "white"
		dictionary := Dictionary{key: color}
		newKey := "shape"
		newColor := "blue"

		err := dictionary.Update(newKey, newColor)

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete a key/value pair ina  dictionary", func(t *testing.T) {
		key := "color"
		color := "white"
		dictionary := Dictionary{key: color}
		dictionary.Delete(key)

		_, exist := dictionary.Search(key)

		if exist != ErrNotFound {
			t.Errorf("did not find any error when we should've")
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q...want %q, given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q, wanted %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if got != definition {
		t.Fatal("can't find search word:", err)
	}
	assertStrings(t, got, definition)
}
