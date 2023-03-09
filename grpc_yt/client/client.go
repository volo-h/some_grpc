package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_yt/gen/proto"
	"log"
)

func main(){
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msq: "hello everyone"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.Msq)
}
