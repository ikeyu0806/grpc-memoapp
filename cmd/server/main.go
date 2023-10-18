package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// type server struct{}

// func (s *server) GetMemo(ctx context.Context, in *grpcPkg.GetMemoRequest) (*grpcPkg.GetMemoResponse, error) {
// 	dummyMemo := &grpcPkg.Memo{
// 		Title:       "Sample Memo",
// 		Description: "This is a sample memo.",
// 	}

// 	response := &grpcPkg.GetMemoResponse{
// 		Memo: dummyMemo,
// 	}

// 	return response, nil
// }

// func (s *server) CreateMemo(ctx context.Context, in *grpcPkg.CreateMemoRequest) (*grpcPkg.CreateMemoResponse, error) {
// 	newMemoID := "12345"

// 	response := &grpcPkg.CreateMemoResponse{
// 		Success: true,
// 		Id:      newMemoID,
// 	}

// 	return response, nil
// }

// func (s *server) ListMemos(ctx context.Context, in *grpcPkg.ListMemosRequest) (*grpcPkg.ListMemosResponse, error) {
// 	dummyMemos := []*grpcPkg.Memo{
// 		{
// 			Title:       "Memo 1",
// 			Description: "This is the first memo.",
// 		},
// 		{
// 			Title:       "Memo 2",
// 			Description: "This is the second memo.",
// 		},
// 	}

// 	response := &grpcPkg.ListMemosResponse{
// 		Memos: dummyMemos,
// 	}

// 	return response, nil
// }

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
