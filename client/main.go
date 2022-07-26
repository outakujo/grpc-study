package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func main() {
	addr := flag.String("addr", "localhost:9010", "addr")
	id := flag.String("id", "", "client id")
	recon := flag.Int("recon", 2, "client reconnect time(second)")
	flag.Parse()
	if *id == "" {
		fmt.Println("-id not be empty")
		return
	}
	mycred = MyCred{
		Id:    *id,
		Addr:  *addr,
		Recon: time.Duration(*recon) * time.Second,
	}
	daemon()
}

func daemon() {
	err := register(getClient())
	if err != nil {
		code := status.Code(err)
		if code == codes.Unavailable {
			daemon()
			return
		}
	}
	fmt.Println("register")
	err = recvCtl(getClient())
	if err != nil {
		daemon()
	}
}
