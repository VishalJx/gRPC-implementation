package main

import (
	"bidirectional-grpc/proto/pb"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names : []string{"Priyanshu", "Jay", "Sam"},
	}

	// callSayHello(client) /*Unary*/
	callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	// callHelloBidirectionalStream(client, names)

}