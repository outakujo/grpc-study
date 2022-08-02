### grpc study

#### server

```shell
make build
./ser
```

#### install

```shell
curl "http://localhost:9009/install?id=546" | sh
```

#### stop

```shell
cd agent
./agent.sh stop
```

#### ser -h

```shell
Usage of ./ser:
  -cert string
        ssl cert
  -db-db string
        db database (default "agent")
  -db-host string
        db host (default "127.0.0.1")
  -db-pass string
        db password (default "123456")
  -db-port int
        db port (default 3306)
  -db-user string
        db user (default "root")
  -key string
        ssl key
  -port int
        port (default 9009)
```

#### cli -h

```shell
Usage of ./cli:
  -addr string
        addr (default "localhost:9010")
  -cert string
        ssl cert
  -id string
        client id
  -recon int
        client reconnect time(second) (default 2)
  -report int
        client report time(second) (default 10)
```
