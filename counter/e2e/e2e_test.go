// TODO: add more word count tests

package e2e

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/calumball/ori/counter/proto"
)

var client pb.CounterClient

func TestMain(m *testing.M) {
	addr := flag.String("addr", "localhost:8888", "ip:port of counter service")
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to counter service on %v: %v\n", *addr, err)
		os.Exit(1)
	}
	client = pb.NewCounterClient(conn)

	os.Exit(m.Run())
}

func TestGetWordCount(t *testing.T) {
	resp, err := client.GetWordCount(context.Background(), &pb.GetWordCountRequest{Text: "hello\nhello\nhello"})
	if err != nil {
		t.Errorf("failed to get word count: %v\n", err)
	}
	count := resp.GetWordFrequency()["hello"]
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
	{"newline", "\n", 2},
	{"one word, new line", "ori\n", 2},
	{"empty string", "", 0},
	{"many lines", strings.Repeat("\n", 100), 101},
}

func TestGetLineCount(t *testing.T) {
	for _, tt := range linecounttests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.GetLineCount(context.Background(), &pb.GetLineCountRequest{Text: tt.in})
			if err != nil {
				t.Errorf("failed to get line count: %v\n", err)
			}
			count := resp.GetCount()
			if count != tt.out {
				t.Errorf("line count was %d; wanted %d\n", count, tt.out)
			}
		})
	}
}
