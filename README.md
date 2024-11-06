```bash
protoc --proto_path=./pb ./pb/storage.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=./RPC/storage
protoc --proto_path=./pb ./pb/storage.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go_out=./RPC/storage 
```