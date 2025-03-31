package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/pb"
)

type rpcCall struct {
	pb.UnimplementedServiceServer
}

// Call implements the RPC method
func (this *rpcCall) Call(c context.Context, r *pb.Request) (*pb.Response, error) {

	// Unmarshal the request body into a map
	var req map[string]interface{}
	err := json.Unmarshal(r.Body, &req)
	if err != nil {
		log.Printf("Error unmarshaling request body: %v", err)
		return nil, fmt.Errorf("invalid request format")
	}

	log.Printf("Received request: %+v", req)

	// Prepare response data
	res := map[string]interface{}{
		"code":    0,
		"message": "success",
		"data":    map[string]string{"test": "Hello World!"},
	}

	// Marshal the response data
	b, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
		return nil, fmt.Errorf("failed to generate response")
	}

	// Return the response
	return &pb.Response{Response: b}, nil
}

func main() {
	// Listen for incoming connections
	l, err := net.Listen("tcp", ":60001")
	if err != nil {
		log.Fatalf("Failed to listen on port 60001: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the service
	pb.RegisterServiceServer(s, &rpcCall{})

	// Start the server
	log.Println("Server is listening on port 60001...")
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
