package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/proto"
	"log"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("did not close: %v", err)
		}
	}(conn)

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Name: []string{"Naman", "Alice", "BoB"},
	}

	callSayHello(client) // unary api
	callSayHelloServerStream(client, names)
}
