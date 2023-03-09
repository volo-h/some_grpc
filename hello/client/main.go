package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello "hello/proto"
	"log"
)

func main() {
    fmt.Println("Hello client ...")

    opts := grpc.WithInsecure()
    cc, err := grpc.Dial("localhost:50051", opts)
    if err != nil {
	log.Fatal(err)
    }
    defer cc.Close()

    client := hello.NewHelloServiceClient(cc)
    request := &hello.HelloRequest{Name: "vova"}

    resp, _ := client.Hello(context.Background(), request)
    fmt.Printf("Receive response => [%v]", resp.Greeting)
}
