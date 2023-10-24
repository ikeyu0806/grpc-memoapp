package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	memopb "grpc-memoapp/pkg/grpc"
)

type memoServer struct {
	memopb.UnimplementedMemoAPIServer
}

func NewMemoServer() *memoServer {
	return &memoServer{}
}

func (s *memoServer) GetMemo(ctx context.Context, req *memopb.GetMemoRequest) (*memopb.GetMemoResponse, error) {
	dummyMemo := &memopb.Memo{
		Title:       "Sample Memo",
		Description: "This is a sample memo.",
	}

	response := &memopb.GetMemoResponse{
		Memo: dummyMemo,
	}

	return response, nil
}

func (s *memoServer) CreateMemo(ctx context.Context, req *memopb.CreateMemoRequest) (*memopb.CreateMemoResponse, error) {
	os.Remove("./grpc_memoapp.db")

	db, err := sql.Open("sqlite3", "./grpc_memoapp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS memos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Println(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into memos(title, description) values('hoge', 'huga')")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	newMemoID := "12345"

	response := &memopb.CreateMemoResponse{
		Success: true,
		Id:      newMemoID,
	}

	return response, nil
}

func (s *memoServer) ListMemos(ctx context.Context, req *memopb.ListMemosRequest) (*memopb.ListMemosResponse, error) {
	dummyMemos := []*memopb.Memo{
		{
			Title:       "Memo 1",
			Description: "This is the first memo.",
		},
		{
			Title:       "Memo 2",
			Description: "This is the second memo.",
		},
	}

	response := &memopb.ListMemosResponse{
		Memos: dummyMemos,
	}

	return response, nil
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	memopb.RegisterMemoAPIServer(s, NewMemoServer())

	reflection.Register(s)

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
