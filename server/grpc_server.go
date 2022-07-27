package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-study/pb"
	"net"
	"strconv"
)

type Server struct {
	pb.UnimplementedServiceServer
	Interceptor
}

func (r *Server) Register(_ context.Context, req *pb.Req) (*pb.Resp, error) {
	end := manager.Add(req.Id)
	return &pb.Resp{
		Result: "end:" + end.Id,
	}, nil
}

func (r *Server) Ctl(_ *emptypb.Empty, stream pb.Service_CtlServer) error {
	end, err := r.Auth(stream.Context(), manager)
	if err != nil {
		return err
	}
	defer func() {
		manager.Del(end.Id)
		fmt.Printf("del id=%s\n", end.Id)
	}()
	for {
		err = stream.Send(<-end.Ch)
		if err != nil {
			return err
		}
	}
}

func (r *Server) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (r *Server) Report(stream pb.Service_ReportServer) error {
	end, err := r.Auth(stream.Context(), manager)
	if err != nil {
		return err
	}
	defer func() {
		manager.Del(end.Id)
		fmt.Printf("del id=%s\n", end.Id)
	}()
	for {
		var sysInfo *pb.SysInfo
		sysInfo, err = stream.Recv()
		if err != nil {
			return err
		}
		fmt.Println(end.Id, sysInfo)
	}
}

func runGrpcServer(port int) {
	var ser Server
	server := grpc.NewServer()
	pb.RegisterServiceServer(server, &ser)
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port+1))
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
