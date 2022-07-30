prot:
	protoc -I proto --go_out=.  --go-grpc_out=. \
	proto/svc.proto
ser:
	go build -o ser ./server
cli:
	mkdir -p static
	go build -o ./static/cli ./client
script:
	 go-bindata -pkg control -o control/script.go control/script
build: ser cli