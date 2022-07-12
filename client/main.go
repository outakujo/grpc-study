package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-study/pb"
	"time"
)

func main() {
	conn, err := grpc.Dial(":9009", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.Call(ctx, &pb.Req{Param: "定时发"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Result)
}
