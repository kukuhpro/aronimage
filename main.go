package main

import (
	agrpc "aronimage/grpc"
	"log"
	"net"
	"qasirworker"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "aronimage/protocols"
)

const (
	port      = ":50051"
	MaxWorker = 5
	MaxQueue  = 10000
)

var server *agrpc.Server

// intiate function
func init() {
	server = agrpc.NewServer()
	qasirworker.NewDispatcher(MaxWorker, MaxQueue).Run()
}

// main function
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		color.Red("failed to listen: %v", err)
	} else {
		color.Green("Running GRPC Server on port " + port)
	}

	s := grpc.NewServer()
	pb.RegisterImageServer(s, server)
	reflection.Register(s)
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
