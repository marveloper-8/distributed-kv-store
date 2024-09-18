package main

import (
	"context"
	"log"
	"time"

	pb "distributed-kv-store/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("path/to/certifile", "")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewKeyValueStoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.Put(ctx, &pb.PutRequest{Key: "foo", Value: "bar"})
	if err != nil {
		log.Fatalf("could not put: %v", err)
	}

	getResp, err := c.Get(ctx, &pb.GetRequest{Key: "foo"})
	if err != nil {
		log.Fatalf("could not get %v", err)
	}
	log.Printf("Get result: %s", getResp.Value)
}
