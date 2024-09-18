package main

import (
	"log"
	"net"

	pb "distributed-kv-store/api/proto"
	"distributed-kv-store/pkg/node_communication"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterKeyValueStoreServer(grpcServer, node_communication.NewServer())

	log.Println("Server is running on port :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}