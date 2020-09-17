package main

import (
	"context"
	"github.com/d3z41k/grpc-basic/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct{}

func main() {
	lister, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(lister); err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.ResponseInt, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.ResponseInt{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.ResponseInt, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.ResponseInt{Result: result}, nil
}

func (s *server) Division(ctx context.Context, request *proto.Request) (*proto.ResponseFloat, error) {
	a, b := request.GetA(), request.GetB()

	result := float32(a) / float32(b)

	return &proto.ResponseFloat{Result: result}, nil
}
