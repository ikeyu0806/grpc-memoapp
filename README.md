
## 環境構築
```bash
docker-compose up
```

### migration
```bash
docker-compose run grpc-memoapp go run cmd/migrate/main.go
```

## protoファイルからの自動コード生成
参考: https://grpc.io/docs/languages/go/quickstart/

```bash
brew install protobuf

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

```bash
cd proto
protoc --go_out=../pkg/grpc --go_opt=paths=source_relative memoservice.proto --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative memoservice.proto
```

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/memoservice.proto
```

## grpcurlでの動作検証
```bash
brew install grpcurl

# サービス一覧確認
grpcurl -plaintext localhost:3333 list
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
service.MemoAPI

# メソッド一覧確認
grpcurl -plaintext localhost:3333 list service.MemoAPI
service.MemoAPI.CreateMemo
service.MemoAPI.GetMemo
service.MemoAPI.ListMemos

grpcurl -plaintext -d '{"id": "1"}' localhost:3333 service.MemoAPI.GetMemo
grpcurl -plaintext localhost:3333 service.MemoAPI.ListMemos
grpcurl -plaintext -d '{
  "memo": {
    "title": "titleDemo",
    "description": "descriptionDemo"
  }
}' localhost:3333 service.MemoAPI.CreateMemo
grpcurl -plaintext -d '{
  "memo": {
    "id": "1",
    "title": "updateDemo",
    "description": "updateDemo"
  }
}' localhost:3333 service.MemoAPI.UpdateMemo
grpcurl -plaintext -d '{"id": "1"}' localhost:3333 service.MemoAPI.DeleteMemo
```

# Goのクライアントコマンド実行
```bash
# 最後の<1>はID指定
docker compose run grpc-memoapp go run cmd/client/main.go 1
```

# sqlite操作
```
docker exec -it grpc-memoapp_grpc-memoapp_1 bash
sqlite3 grpc_memoapp.db
# テーブル一覧
.tables
```
