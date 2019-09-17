//TODO: add more word count tests

package server

import (
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	counts := countWords("hello\nhello\nhello", true)
	count := counts["hello"]
	if count != 3 {
		t.Errorf("word count was %d; wanted 3\n", count)
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
