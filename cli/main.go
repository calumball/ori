/*
A command-line interface for interacting with the counter gRPC server.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"

	"google.golang.org/grpc"

	pb "github.com/calumball/ori/counter/proto"
)

const (
	defaultAddress = "localhost:8888"
)

func main() {
	addr := flag.String("addr", defaultAddress, "ip:port of counter service")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 || !(args[0] == "words" || args[0] == "lines") {
		fmt.Fprintf(os.Stderr, "usage: counter-cli --addr=%v [command: words or lines] [text]\n", defaultAddress)
		os.Exit(1)
	}

	cmd := args[0]
	text := args[1]

	client, err := startClient(*addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to counter service on %v: %v\n", *addr, err)
		os.Exit(1)
	}

	err = handleCommand(client, cmd, text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not handle command: %v\n", err)
		os.Exit(1)
	}
}

func startClient(addr string) (pb.CounterClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return pb.NewCounterClient(conn), nil
}

func handleCommand(client pb.CounterClient, cmd, text string) error {
	switch cmd {
	case "words":
		req := pb.GetWordCountRequest{Text: text}
		resp, err := client.GetWordCount(context.Background(), &req)
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not get word count")
			return err
		}
		sortAndPrintWordCount(resp.GetWordFrequency())
	case "lines":
		req := pb.GetLineCountRequest{Text: text}
		resp, err := client.GetLineCount(context.Background(), &req)
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not get line count")
			return err
		}
		fmt.Fprintf(os.Stdout, "line count: %d\n", resp.GetCount())
	}
	return nil
}

func sortAndPrintWordCount(freq map[string]uint64) {
	n := map[int][]string{}
	var a []int

	for word, count := range freq {
		count := int(count)
		n[count] = append(n[count], word)
	}

	for count := range n {
		a = append(a, count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	fmt.Fprintln(os.Stdout, "word counts:")
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Fprintf(os.Stdout, "%s, %d\n", s, k)
		}
	}
}
