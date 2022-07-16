package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-study/control"
	"grpc-study/pb"
	"time"
)

type MyCred struct {
	Id string
}

func (m *MyCred) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{"id": m.Id}, nil
}

func (m *MyCred) RequireTransportSecurity() bool {
	return false
}

func main() {
	addr := flag.String("addr", "localhost:9010", "addr")
	flag.Parse()
	credentials := grpc.WithPerRPCCredentials(&MyCred{Id: "123"})
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.
		NewCredentials()), credentials)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := c.Call(ctx, &pb.Req{Param: "test"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Result)
	ctl, err := c.Ctl(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	for {
		cmd, err := ctl.Recv()
		if err != nil {
			panic(err)
		}
		result, err := control.Exe(nil, cmd.Name, cmd.Arg...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print(result)
	}
}
