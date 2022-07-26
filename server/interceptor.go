package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var NotAuthError = errors.New("not auth")
var NotExistError = errors.New("not exist")
var NotAuthRpcError = status.Error(NotAuth, NotAuthError.Error())
var NotExistRpcError = status.Error(NotExist, NotExistError.Error())

const (
	NotAuth codes.Code = 1001 + iota
	NotExist
)

type Interceptor struct {
}

func (r Interceptor) Auth(ctx context.Context, man *Manager) (*End, error) {
	ic, _ := metadata.FromIncomingContext(ctx)
	ids := ic.Get("id")
	if len(ids) == 0 {
		return nil, NotAuthRpcError
	}
	s := ids[0]
	fmt.Printf("auth id=%s\n", s)
	end := man.Get(s)
	if end == nil {
		return nil, NotExistRpcError
	}
	return end, nil
}
