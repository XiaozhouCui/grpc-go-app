## Install gRPC

- Go to https://grpc.io/ and follow the steps for Go
- Install **Protocol Buffer Compiler**: `brew install protobuf`
- Verify installation: `protoc --version`
- Install Go plugins for protocol compiler
- `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
- `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
- Update PATH so that the protoc compiler can find the plugins

```
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
```

- In VS Code, install plugin `vscode-proto3`

## Initialize project folder

- Create and go to folder `grpc-go-app`
- Run `go mod init github.com/XiaozhouCui/grpc-go-app` to init go project
- Run `go mod tidy`
- Run `git init` to init repo

## Create dummy protocol buffer Go files from proto

- Create file _greet/proto/dummy.proto_
- Generate Go code from proto file
- Run `protoc -Igreet/proto --go_out=. --go_opt=module=github.com/XiaozhouCui/grpc-go-app --go-grpc_out=. --go-grpc_opt=module=github.com/XiaozhouCui/grpc-go-app greet/proto/dummy.proto`
- To make CLI easier, create a Makefile and run `make greet`, will generate the same files
- To remove the generated files, run `make clean_greet`
- For help about the Makefile, run `make help`
- To show system information, run `make about`

## Setup Greet server

- Clean up the dummy files in greet: `make clean_greet`
- Create _greet.proto_, then run `make greet` to generate protocol buffer go files
- Create _./greet/server/main.go_ to create net server
- Run `go mod tidy` to install all missing dependencies, generating _go.sum_
- Run `make greet` again to build the server bin file
- To start the server, run `./bin/greet/server`, should see "Listening on 0.0.0.0:50051"

## Setup Greet client

- Keep the server running
- Create _./greet/client/main.go_ to create client
- Run `make greet` again to build the client bin file
- Run `./bin/greet/client` to test client, should see no error

## Unary API server implementation

- Update _greet/server/main.go_ to register GreetServiceServer
- Create _greet/server/greet.go_ to implement rpc endpoints
