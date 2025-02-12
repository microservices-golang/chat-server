package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/microservices-golang/chat-server/pkg/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	chat.UnimplementedChatServiceServer
}

var grpcPort = 50052

// protoc  --go_out=./pkg/ --go-grpc_out=./pkg api/chat_v1/chat.proto
func (s *server) CreateChat(_ context.Context, req *chat.CreateChatRequest) (*chat.CreateChatResponse, error) {
	fmt.Println(req.Usernames)
	return &chat.CreateChatResponse{Id: 1}, nil
}

func (s *server) DeleteChat(_ context.Context, req *chat.DeleteChatRequest) (*emptypb.Empty, error) {
	fmt.Println(req.Id)
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, req *chat.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Printf("SendMessage request: from=%s, text=%s, timestamp=%v\n", req.From, req.Text, req.Timestamp)
	return &emptypb.Empty{}, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat.RegisterChatServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
