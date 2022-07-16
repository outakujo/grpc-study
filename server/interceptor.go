package main

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

var NotAuthError = errors.New("not auth")

type Interceptor struct {
}

func (r Interceptor) Auth(ctx context.Context, man *Manager) (*End, error) {
	ic, _ := metadata.FromIncomingContext(ctx)
	ids := ic.Get("id")
	if len(ids) == 0 {
		return nil, NotAuthError
	}
	s := ids[0]
	end := man.Add(s)
	return end, nil
}
