package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-study/pb"
	"net"
	"net/http"
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
	}()
	for {
		err = stream.Send(<-end.Ch)
		if err != nil {
			return err
		}
	}
}

var manager = NewManager()

func main() {
	port := flag.Int("port", 9009, "port")
	flag.Parse()
	var ser Server
	engine := gin.New()
	engine.POST("ctl", func(c *gin.Context) {
		var da = struct {
			Id   string
			Name string
			Arg  []string
		}{}
		err := c.ShouldBind(&da)
		if err != nil {
			c.JSON(http.StatusOK, err.Error())
			return
		}
		if da.Id == "" {
			c.JSON(http.StatusOK, "id not be empty")
			return
		}
		if da.Name == "" {
			c.JSON(http.StatusOK, "name not be empty")
			return
		}
		end := manager.Get(da.Id)
		if end == nil {
			c.JSON(http.StatusOK, NotExistError.Error())
			return
		}
		end.Ch <- &pb.Cmd{Name: da.Name, Arg: da.Arg}
		c.JSON(http.StatusOK, "ok")
	})
	go func() {
		err := engine.Run(":" + strconv.Itoa(*port))
		if err != nil {
			panic(err)
		}
	}()
	server := grpc.NewServer()
	pb.RegisterServiceServer(server, &ser)
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(*port+1))
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
