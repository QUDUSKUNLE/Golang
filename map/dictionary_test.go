package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		ok, _ := dictionary.Search("test")
		want := "this is just a test"
		AssertStrings(t, ok, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}
		AssertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		AssertError(t, nil, err)
		AssertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		AssertError(t, err, ErrWordExists)
		AssertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
	
		AssertError(t, err, nil)
		AssertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("update new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
	
		AssertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delate word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}
		dictionary.Delete(word)
	
		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	AssertStrings(t, got, definition)
}
