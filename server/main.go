package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-study/pb"
	"net"
	"net/http"
)

type Server struct {
	pb.UnimplementedServiceServer
}

func (r *Server) Call(_ context.Context, req *pb.Req) (*pb.Resp, error) {
	return &pb.Resp{
		Result: req.Param + "ddd",
	}, nil
}

func (r *Server) Ctl(_ *emptypb.Empty, steam pb.Service_CtlServer) error {
	end := steam.Context().Value("end").(*End)
	for {
		err := steam.Send(<-end.Ch)
		if err != nil {
			return err
		}
	}
}

func main() {
	var ser Server
	manager := NewManager()
	interceptor := grpc.UnaryInterceptor(func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ic, _ := metadata.FromIncomingContext(ctx)
		ids := ic.Get("id")
		if len(ids) == 0 {
			return nil, errors.New("not auth")
		}
		id := ids[0]
		end := manager.Add(id)
		ctx = context.WithValue(ctx, "end", end)
		return handler(ctx, req)
	})
	engine := gin.New()
	engine.GET("ctl", func(c *gin.Context) {
		query := c.Query("id")
		if query == "" {
			c.JSON(http.StatusOK, "id empty")
			return
		}
		end := manager.Add(query)
		end.Ch <- &pb.Cmd{Name: "go", Arg: []string{"version"}}
		c.JSON(http.StatusOK, "ok")
	})
	go func() {
		engine.Run()
	}()
	server := grpc.NewServer(interceptor)
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
