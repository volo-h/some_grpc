package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pb "grpc_yt/gen/proto"
	"log"
	"net"
	"net/http"
)

type testApiServer struct {
  pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error)	 {
	return req, nil
}

func main() {
	go func() {
		mux := runtime.NewServeMux()
		pb.RegisterTestApiHandlerServer(context.Background(), mux, &testApiServer{})
		log.Fatalln(http.ListenAndServe("localhost:8081", mux))
	}()

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
  	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Println(err)
	}
}
