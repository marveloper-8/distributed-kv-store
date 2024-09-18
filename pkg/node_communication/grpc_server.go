package node_communication

import (
	"context"
	pb "distributed-kv-store/api/proto"
	"log"
	"sync"
)

type Server struct {
	pb.UnimplementedKeyValueStoreServer
	mu sync.RWMutex
	store map[string]string
}

func NewServer() *Server {
	return &Server{
		store: make(map[string]string),
	}
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.store[req.Key]
	if !ok {
		return nil, nil
	}
	return &pb.GetResponse{Value: value}, nil
}

func (s *Server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[req.Key] = req.Value
	log.Printf("Stored key-value: %s = %s", req.Key, req.Value)
	return &pb.PutResponse{}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, req.Key)
	log.Printf("Deleted key: %s", req.Key)
	return &pb.DeleteResponse{}, nil
}