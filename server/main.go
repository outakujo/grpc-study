package main

import (
	"flag"
)

func main() {
	port := flag.Int("port", 9009, "port")
	flag.Parse()
	runApiServer(*port)
	runGrpcServer(*port)
}
