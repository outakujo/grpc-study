package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-study/pb"
	"runtime"
	"time"
)

func main() {
	addr := flag.String("addr", "localhost:9010", "addr")
	id := flag.String("id", "", "client id")
	recon := flag.Int("recon", 2, "client reconnect time(second)")
	repor := flag.Int("report", 10, "client report time(second)")
	flag.Parse()
	if *id == "" {
		fmt.Println("-id not be empty")
		return
	}
	mycred = MyCred{
		Id:    *id,
		Addr:  *addr,
		Recon: time.Duration(*recon) * time.Second,
		Ch:    make(chan *pb.SysInfo),
	}
	daemon()
	for {
		info := &pb.SysInfo{Os: runtime.GOOS, Arch: runtime.GOARCH}
		mycred.Ch <- info
		time.Sleep(time.Duration(*repor) * time.Second)
	}
}

func daemon() {
	err := register(getClient())
	if isUnavailable(err) {
		daemon()
		return
	}
	fmt.Println("register")
	go func() {
		err = recvCtl(getClient())
		if isUnavailable(err) {
			daemon()
		}
	}()
	go func() {
		err = report(getClient(), mycred.Ch)
		if isUnavailable(err) {
			daemon()
		}
	}()
}

func isUnavailable(err error) bool {
	if err != nil {
		code := status.Code(err)
		if code == codes.Unavailable {
			return true
		}
	}
	return false
}
