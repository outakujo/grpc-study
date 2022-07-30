package main

import (
	"github.com/gin-gonic/gin"
	"grpc-study/pb"
	"net/http"
	"strconv"
	"strings"
)

var installSh = `#!/bin/bash
home=%s
if [ -z "$home" ]; then
    home=agent
fi
id=%s
server=%s
grpcAddr=%s
download="http://${server}/static/cli"
mkdir -p $home
cd $home || exit 1
curl -O -k $download
chmod +x ./cli

cat > agent.sh << EOF
#!/bin/bash
id=$id
grpcAddr=$grpcAddr

start() {
	echo start	
	nohup ./cli -id \$id -addr \$grpcAddr >log 2>&1
}

stop() {
	echo stop
}

case \$1 in
start)
  start
  ;;
stop)
  stop
  ;;
*)
  echo "Usage: \$0 {start|stop}"
  ;;
esac
EOF

if [ $? -eq 0 ]; then
	chmod +x agent.sh
    ./agent.sh start
else
	echo failed
fi

`

func runApiServer(port int) {
	engine := gin.New()
	engine.Static("static", "static")
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
	engine.GET("install", func(c *gin.Context) {
		home := c.Query("home")
		id := c.Query("id")
		if id == "" {
			c.String(http.StatusOK, "id not be empty")
			return
		}
		host := c.Request.Host
		split := strings.Split(host, ":")
		grpcAddr := ""
		if len(split) == 2 {
			grpcAddr = split[0] + ":" + strconv.Itoa(port+1)
		}
		c.String(http.StatusOK, installSh, home, id, host, grpcAddr)
	})
	go func() {
		err := engine.Run(":" + strconv.Itoa(port))
		if err != nil {
			panic(err)
		}
	}()
}
