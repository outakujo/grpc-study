package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
		var ev *pb.Event
		ev, err = stream.Recv()
		if err != nil {
			return err
		}
		switch ev.Detail.(type) {
		case *pb.Event_SysInfo_:
			sysInfo := ev.GetSysInfo()
			info := SysInfo{
				Id:   end.Id,
				Os:   sysInfo.Os,
				Arch: sysInfo.Arch,
			}
			err = service.SaveSysInfo(&info)
		case *pb.Event_ScriptResult_:
			result := ev.GetScriptResult()
			info := ScriptResult{
				Id:     end.Id,
				Code:   int(result.Code),
				Stdout: result.Stdout,
			}
			err = service.SaveScriptResult(&info)
		default:
			fmt.Println("unknow event type ")
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}

func runGrpcServer(port int, cert, key string) {
	var ser Server
	options := make([]grpc.ServerOption, 0)
	if cert != "" && key != "" {
		file, err := credentials.NewServerTLSFromFile(cert, key)
		if err != nil {
			panic(err)
		}
		options = append(options, grpc.Creds(file))
	}
	server := grpc.NewServer(options...)
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
