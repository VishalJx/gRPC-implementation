package main

import (
	"bidirectional-grpc/proto/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main()  {
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("Failed to start SERVER %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server Started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to Start: %v", err)
	}
}


