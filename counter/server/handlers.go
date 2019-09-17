package server

import (
	"context"

	pb "github.com/calumball/ori/counter/proto"
)

// CounterServer handles counting requests.
type CounterServer struct{}

// GetWordCount counts words in a text.
func (CounterServer) GetWordCount(context context.Context, req *pb.GetWordCountRequest) (*pb.WordCount, error) {
	count := countWords(req.Text, req.RespectCaps)
	return &pb.WordCount{WordFrequency: count}, nil
}

//GetLineCount counts lines in a text.
func (CounterServer) GetLineCount(context context.Context, req *pb.GetLineCountRequest) (*pb.LineCount, error) {
	count := countLines(req.Text)
	return &pb.LineCount{Count: count}, nil
}
