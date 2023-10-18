
## 環境構築
```bash
docker-compose up
```

## protoファイルからの自動コード生成
参考: https://grpc.io/docs/languages/go/quickstart/

```bash
brew install protobuf

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

```bash
protoc --go_out=../pkg/grpc --go_opt=paths=source_relative memoservice.proto --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative memoservice.proto
```

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/memoservice.proto
```
