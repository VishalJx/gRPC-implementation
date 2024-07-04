package main

import (
	"bidirectional-grpc/proto/pb"
	"context"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam)(*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}