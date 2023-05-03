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
- Create _greet/server/greet.go_ to implement Greet RPC endpoint

## Unary API client implementation

- Update _greet/client/main.go_ to create an instance of greet service client
- Create _greet/client/greet.go_ to call Greet RPC endpoint
- Run `make greet` to re-build binary files
- Run `./bin/greet/server` to start gRPC server
- Stert a new terminal, run `./bin/greet/client` to call gRPC server
- Should see "Response from Greet RPC: Hello Joe"

## Setup SSL for gRPC

- create file _./ssl/ssl.sh_
- Go to ssl folder, and run `chmod +x ssl.sh` to make file executable
- Then run `./ssl.sh` to generate certificates and keys
- Server: Update _./greet/server/main.go_, pass SSL creds options into `grpc.NewServer()` to apply SSL
- Client: Update _./greet/client/main.go_, pass SSL creds options into `grpc.Dial()` to apply SSL
- Run `make greet`, `./bin/greet/server` and then `./bin/greet/client` to test, everything should work as before

## gRPC Reflection (find server functions and services)

- Evans can call gRPC endpoints without client
- Install Evans CLI: `brew tap ktr0731/evans` then `brew install evans`
- Validate installation by running `evans`
- Update _./calculator/server/main.go_ to register reflection
- Run `make calculator` then start server `./bin/calculator/server`
- Start evans CLI `evans --host localhost --port 50051 --reflection repl`
- Inside evans CLI, run `show package` to show calculator package
- Enter `package calculator` to select the calculator package
- Enter `show message`, to see messages
- Enter `show service`, to see services (Sum, Primes, Avg, Max, Sqrt)
- Select calculator service `service CalculatorService`
- Unary: Inside CalculatorService, run `call Sum`, then erter `first_number` and `second_number`, then it should return the sumResult
- Server streamnig: run `call Primes` then enter `number`, it should return the response stream
- Client streaming: Run `call Avg`, enter multiple inputs, then enter `ctrl + d` to stop stream
- Bi-directional: run `call Max`, enter multiple inputs, then enter `ctrl + d` to stop stream
- Error handling: `call Sqrt`, enter a negative number to see error response
- Enter `exit` to quite Evans CLI

## Setup Blog project

- Create _./blog/docker-compose.yaml_
- Go to _./blog_, then run `docker-comopse up`, this will run mongo and mongo-express
- Validate the services by opening `localhost:8081` in borwser
- Install mongo-go-driver `go get go.mongodb.org/mongo-driver/mongo`
- Run `go mod tidy` to update dependencies
