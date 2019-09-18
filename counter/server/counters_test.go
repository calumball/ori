package server

import (
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	counts := countWords("hello\nhello\nhello", false)
	count := counts["hello"]
	if count != 3 {
		t.Errorf("word count was %d; wanted 3\n", count)
	}
}

func TestCountMultipleWords(t *testing.T) {
	counts := countWords("hello\nhello\nhello\nthere there", false)
	helloCount := counts["hello"]
	if helloCount != 3 {
		t.Errorf("word count was %d; wanted 3\n", helloCount)
	}
	thereCount := counts["there"]
	if thereCount != 2 {
		t.Errorf("word count was %d; wanted 2\n", thereCount)
	}
}

func TestCountWordsEmptyString(t *testing.T) {
	counts := countWords("", false)
	if len(counts) != 0 {
		t.Errorf("word count was not empty: %v\n", counts)
	}
}

func TestCountWordsWithCaps(t *testing.T) {
	counts := countWords("hello... HELLO!", true)
	capCount := counts["HELLO"]
	lowerCount := counts["hello"]
	if capCount != 1 {
		t.Errorf("word count was %d; expected 1\n", capCount)
	}
	if lowerCount != 1 {
		t.Errorf("word count was %d; expected 1\n", lowerCount)
	}
}

func TestCountWordsWithoutCaps(t *testing.T) {
	counts := countWords("hello... HELLO!", false)
	count := counts["hello"]
	if count != 2 {
		t.Errorf("word count was %d; expected 2\n", count)
	}
}

var linecounttests = []struct {
	name string
	in   string
	out  uint64
}{
	{"one word", "ori", 1},
	{"empty string", "", 0},
	{"newline", "\n", 2},
	{"one word, new line", "ori\n", 2},
	{"many lines", strings.Repeat("\n", 100), 101},
}

func TestCountLines(t *testing.T) {
	for _, tt := range linecounttests {
		t.Run(tt.name, func(t *testing.T) {
			count := countLines(tt.in)
			if count != tt.out {
				t.Errorf("line count was %d; wanted %d\n", count, tt.out)
			}
		})
	}
}
