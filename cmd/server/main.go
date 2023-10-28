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

	insertSqlStmt := "INSERT INTO memos(title, description) VALUES(?, ?);"
	_, err = db.Exec(insertSqlStmt, req.Memo.Title, req.Memo.Description)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	newMemoID := "12345"

	response := &memopb.CreateMemoResponse{
		Success: true,
		Id:      newMemoID,
	}

	log.Println("success CreateMemo")
	return response, nil
}

func (s *memoServer) ListMemos(ctx context.Context, req *memopb.ListMemosRequest) (*memopb.ListMemosResponse, error) {
	db, err := sql.Open("sqlite3", "./grpc_memoapp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	SELECT * FROM memos;
	`
	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var grpcMemos []*memopb.Memo
	for rows.Next() {
		var id int
		var title, description string
		err := rows.Scan(&id, &title, &description)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Printf("title: %v", title)
		log.Printf("description: %v", description)
		grpcMemo := &memopb.Memo{
			Title:       title,
			Description: description,
		}
		grpcMemos = append(grpcMemos, grpcMemo)
	}

	response := &memopb.ListMemosResponse{
		Memos: grpcMemos,
	}

	log.Println("success ListMemos")

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
