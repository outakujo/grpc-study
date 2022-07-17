package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
)

var NotAuthError = errors.New("not auth")
var NotExistError = errors.New("not exist")

type Interceptor struct {
}

func (r Interceptor) Auth(ctx context.Context, man *Manager) (*End, error) {
	ic, _ := metadata.FromIncomingContext(ctx)
	ids := ic.Get("id")
	if len(ids) == 0 {
		return nil, NotAuthError
	}
	s := ids[0]
	fmt.Printf("ctl id=%s\n", s)
	end := man.Get(s)
	if end == nil {
		return nil, NotExistError
	}
	return end, nil
}
