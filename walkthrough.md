## Install gRPC

- Go to https://grpc.io/ and follow the steps for Go
- Install **Protocol Buffer Compiler**: `brew install protobuf`
- Verify installation: `protoc --version`
- Install Go plugins for protocol compiler
- `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
- `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
- Update PATH so that the protoc compiler can find the plugins

```
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Initialize project

- Run `go mod init github.com/XiaozhouCui/grpc-go-app` to create go project
- Run `go mod tidy`
- Run `git init` to init repo
