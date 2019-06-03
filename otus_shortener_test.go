package otus_shortener

import (
	"math/rand"
	"testing"
)

var samples = []map[string]string{
	{
		"url":  "https://google.com",
		"hash": "4236168800",
	},
	{
		"url":  "http://ya.ru",
		"hash": "1592502014",
	},
	{
		"url":  "http://ya.ru",
		"hash": "15925020141",
	},
	{
		"url":  "",
		"hash": "2166136261",
	},
}

func TestShorten(t *testing.T) {
	s := Storage{}

	rand.Seed(1)

	for _, sample := range samples {
		actual := s.Shorten(sample["url"])
		expected := sample["hash"]

		if actual != expected {
			t.Errorf("Shorten returns incorrect values!\n "+
				"got:      %s \n "+
				"expected: %s",
				actual, expected)
		}
	}
}

func TestResolve(t *testing.T) {
	s := Storage{
		"4236168800":  "https://google.com",
		"1592502014":  "http://ya.ru",
		"15925020141": "http://ya.ru",
		"2166136261":  "",
	}

	for _, sample := range samples {
		actual := s.Resolve(sample["hash"])
		expected := sample["url"]

		if actual != expected {
			t.Errorf("Resolve returns incorrect values!\n "+
				"got:      %s \n "+
				"expected: %s",
				actual, expected)
		}
	}
}
