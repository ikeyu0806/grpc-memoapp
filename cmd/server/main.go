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
	db, err := sql.Open("sqlite3", "./grpc_memoapp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := req.Id

	row := db.QueryRow("SELECT title, description FROM memos WHERE id = ?", id)
	memo := &memopb.Memo{}

	err = row.Scan(&memo.Title, &memo.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Memo not found")
			return nil, err
		}
		log.Println("Error fetching memo")
		return nil, err
	}

	response := &memopb.GetMemoResponse{
		Memo: memo,
	}

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

func (s *memoServer) CreateMemo(ctx context.Context, req *memopb.CreateMemoRequest) (*memopb.CreateMemoResponse, error) {
	db, err := sql.Open("sqlite3", "./grpc_memoapp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	insertSqlStmt := "INSERT INTO memos(title, description) VALUES(?, ?);"
	_, err = db.Exec(insertSqlStmt, req.Memo.Title, req.Memo.Description)

	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	var lastInsertID string
	// SELECT last_insert_rowid()で最後のIDが取得できないのでこの実装。今回はgRPCの実験なので深追いしない
	err = db.QueryRow("SELECT id FROM memos ORDER BY id DESC LIMIT 1").Scan(&lastInsertID)
	if err != nil {
		log.Println(err)
	}

	response := &memopb.CreateMemoResponse{
		Success: true,
		Id:      lastInsertID,
	}

	log.Println("success CreateMemo")
	return response, nil
}

func (s *memoServer) UpdateMemo(ctx context.Context, req *memopb.UpdateMemoRequest) (*memopb.UpdateMemoResponse, error) {
	db, err := sql.Open("sqlite3", "./grpc_memoapp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	updateSqlStmt := "UPDATE memos SET title = ?, description = ? where id = ?;"
	_, err = db.Exec(updateSqlStmt, req.Memo.Title, req.Memo.Description, req.Memo.Id)

	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	response := &memopb.UpdateMemoResponse{
		Success: true,
		Id:      req.Memo.Id,
	}

	log.Println("success UpdateMemo")
	return response, nil
}

func (s *memoServer) DeleteMemo(ctx context.Context, req *memopb.DeleteMemoRequest) (*memopb.DeleteMemoResponse, error) {
	db, err := sql.Open("sqlite3", "./grpc_memoapp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := req.Id

	_, err = db.Exec("DELETE FROM memos WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	response := &memopb.DeleteMemoResponse{
		Id: id,
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
