package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/calumball/ori/counter/proto"
	server "github.com/calumball/ori/counter/server"
)

const (
	defaultPort = "8888"
)

func main() {
	port := os.Getenv("COUNTER_PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("starting counter server on %v\n", port)
	startServer(port)
}

func startServer(port string) {
	address := fmt.Sprintf(":%v", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("started listening")

	var opts []grpc.ServerOption
	gRPCServer := grpc.NewServer(opts...)
	pb.RegisterCounterServer(gRPCServer, server.CounterServer{})
	log.Println("registered server")

	err = gRPCServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
