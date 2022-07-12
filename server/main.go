package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-study/pb"
	"net"
)

type Server struct {
	pb.UnimplementedServiceServer
}

func (r *Server) Call(_ context.Context, req *pb.Req) (*pb.Resp, error) {
	return &pb.Resp{
		Result: req.Param + "ddd",
	}, nil
}

func main() {
	var ser Server
	server := grpc.NewServer()
	pb.RegisterServiceServer(server, &ser)
	listen, err := net.Listen("tcp", ":9009")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
