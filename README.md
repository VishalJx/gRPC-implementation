# gRPC Implementation Server & Client

This project is a gRPC implementation server that supports client streaming, server streaming, and bidirectional streaming. It provides examples of how to implement and use these features in a gRPC service using Golang.

## Features

- **Client Streaming**: The client sends a stream of messages to the server. Once the client has finished writing the messages, it waits for the server to read them and return a response.
- **Server Streaming**: The client sends a single request to the server and receives a stream of messages in response.
- **Bidirectional Streaming**: Both the client and the server send a stream of messages to each other.

## Requisites

#### 1. Protoc Compiler
To download:

**Linux:**
```bash
$ apt install -y protobuf-compiler
$ protoc --version
```

**Mac:**
```bash
$ brew install protobuf
$ protoc --version
```

**Windows:**
1. Go to this repo: https://github.com/protocolbuffers/protobuf/releases
2. Download the windows compatible latest zip file
3. Unzip the file into `C:/`
4. Copy the path e.g., `C:\protoc-27.2-win64\bin` to set it as environment variables
5. Also add the path to Environment Variables

**To check version:**
```bash
protoc --version
```

#### 2. Golang compiler
- Since I'll be using Golang to implement gRPC 
- However we know that gRPC is polyglot and supports multiple languages and the process to create gRPC each time would be similar.

## How to run.
#### Step 1: Generating protobuf(already generated in this repo)
- In the root, run:
```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```
- This will generate two protobuf files inside a `pb` folder

#### Step 2: Select the desired mode
```bash
cd ./client
```
- open `main.go` file and comment out the non required modes. 


#### Step 3: Run the server first
```bash
cd ./client
go run .
```

#### Step 4: Run the client: new terminal
```bash
cd ./server
go run .
```
