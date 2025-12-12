package hades

import (
	"testing"
)

func TestHasRussianChars(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"empty", "", false},
		{"english only", "hello world", false},
		{"russian letters", "привет", true},
		{"mixed", "hello привет", true},
		{"cyrillic supplement", "ҐЄІЇ", true},
		{"non-cyrillic unicode", "こんにちは", false},
		{"numbers and punctuation", "12345!@#$%", false},
		{"cyrillic with numbers", "тест123", true},
		{"cyrillic with punctuation", "тест!", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := HasRussianChars(c.in)
			if got != c.want {
				t.Fatalf("HasRussianChars(%q) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}

func TestCountURLs(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want int
	}{
		{"no urls", "hello world", 0},
		{"one http", "visit http://example.com for info", 1},
		{"one https", "secure https://example.com/path?x=1", 1},
		{"multiple", "links: http://a.com and https://b.com/page and http://c.com", 3},
		{"url-like but no scheme", "www.example.com http:/bad.com", 0},
		{"trailing punctuation", "check http://example.com, and http://foo.com.", 2},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := CountURLs(c.in)
			if got != c.want {
				t.Fatalf("CountURLs(%q) = %d, want %d", c.in, got, c.want)
			}
		})
	}
}
