package main

import (
	"aronimage/grpc"
	"log"
	"net"
	"qasirworker"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "aronimage/protocols"
)

const (
	port      = ":50051"
	MaxWorker = 5
	MaxQueue  = 10000
)

var server *grpc.Server

func init() {
	server = grpc.NewServer()
	qasirworker.NewDispatcher(MaxWorker, MaxQueue).Run()
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Running GRPC Server on port : " + port)
	}

	s := grpc.NewServer()
	pb.RegisterImageServer(s, server)
	reflection.Register(s)
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
