# protobuf demo using Gin

1. Define message formats in a .proto file.
2. Use the protocol buffer compiler.
3. Use the Go protocol buffer API to write and read messages.

## 1. Install protobuf

`brew install protobuf`

`protoc --version` check protoc version

## 2. Defining your protocol format

`awesome.proto`

## 3. Compiling your protocol buffers

install the Go protocol buffers plugin

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

Now run the compiler

`protoc -I=./ --go_out=./protobuf/ ./public/proto/awesome.proto`


## 4. Install protobuf in HTML

Development:
`<script src="//cdn.rawgit.com/dcodeIO/protobuf.js/6.X.X/dist/protobuf.js"></script>`

Production:
`<script src="//cdn.rawgit.com/dcodeIO/protobuf.js/6.X.X/dist/protobuf.min.js"></script>`

## 5. Using the protocol buffers

`go run main.go`

open `http://localhost:8080/`

## References

> [gotutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)

> [gin](https://github.com/gin-gonic/gin)

> [protobuf.js](https://github.com/protobufjs/protobuf.js)