package main

import (
	"github.com/gin-gonic/gin"
	"grpc-study/pb"
	"net/http"
	"strconv"
)

func runApiServer(port int) {
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
		err := engine.Run(":" + strconv.Itoa(port))
		if err != nil {
			panic(err)
		}
	}()
}
