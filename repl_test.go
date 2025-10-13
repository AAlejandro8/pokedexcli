package main

import (
	"testing"
	"time"
	"github.com/AAlejandro8/pokedexcli/internal/pokeapi"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "hello  world",
		expected: []string{"hello", "world"},
	},
	// add more cases here
}
for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("Not the same length!")
		continue
	}
	
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("Words Dont match!")
		}
	}
}

}


func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokeapi.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}