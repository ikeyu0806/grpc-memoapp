
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
protoc --go_out=./ --go-grpc_out=require_unimplemented_servers=false:. ./proto/memoservice.proto
```
