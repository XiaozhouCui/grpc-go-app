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

- In VS Code, install plugin `vscode-proto3`

## Initialize project folder

- Create and go to folder `grpc-go-app`
- Run `go mod init github.com/XiaozhouCui/grpc-go-app` to init go project
- Run `go mod tidy`
- Run `git init` to init repo

## Create Greet project

- Create file _greet/proto/dummy.proto_
- Generate Go code from proto file
- Run `protoc -Igreet/proto --go_out=. --go_opt=module=github.com/XiaozhouCui/grpc-go-app --go-grpc_out=. --go-grpc_opt=module=github.com/XiaozhouCui/grpc-go-app greet/proto/dummy.proto`
- To make CLI easier, create a Makefile and run `make greet`, will do the same
