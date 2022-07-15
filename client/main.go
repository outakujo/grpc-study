package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-study/pb"
	"time"
)

type MyCred struct {
	Id string
}

func (m *MyCred) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"id": m.Id}, nil
}

func (m *MyCred) RequireTransportSecurity() bool {
	return false
}

func main() {
	credentials := grpc.WithPerRPCCredentials(&MyCred{Id: "123"})
	conn, err := grpc.Dial(":9009", grpc.WithTransportCredentials(insecure.
		NewCredentials()), credentials)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := c.Call(ctx, &pb.Req{Param: "定时发"})
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
		fmt.Println(cmd.Name, cmd.Arg)
	}
}
