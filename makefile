prot:
	protoc -I proto --go_out=.  --go-grpc_out=. \
	proto/svc.proto
ser:
	go build -o ser ./server
cli:
	go build -o cli ./client
build: ser cli