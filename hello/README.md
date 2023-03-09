inspired - https://towardsdatascience.com/grpc-in-golang-bb40396eb8b1

terminal1 - go run server/main.go
terminal2 - go run client/main.go

```shell
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto
```