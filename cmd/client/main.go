package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	memopb "grpc-memoapp/pkg/grpc"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main.go <memoID>")
		return
	}

	conn, err := grpc.Dial(
		"grpc-memoapp:8080", // docker-composeのサービス名とポートを指定
		grpc.WithTransportCredentials(insecure.NewCredentials()), // この指定がないとブロックされる
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := memopb.NewMemoAPIClient(conn)

	id := os.Args[1]
	getMemoRequest := &memopb.GetMemoRequest{
		Id: id,
	}
	getMemoResponse, err := client.GetMemo(context.Background(), getMemoRequest)
	if err != nil {
		log.Fatalf("Error calling GetMemo: %v", err)
	}
	fmt.Printf("GetMemo Response: %+v\n", getMemoResponse)
}
