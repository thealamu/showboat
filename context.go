package main

import (
	"context"
	"errors"
)

//context helper methods

type contextKey struct{}

func (s *Server) putUidInContext(baseCtx context.Context, userid string) context.Context {
	return context.WithValue(baseCtx, contextKey{}, userid)
}

func (s *Server) uidFromContext(ctx context.Context) (string, error) {
	val, ok := ctx.Value(contextKey{}).(string)
	if !ok {
		return "", errors.New("could not read uid from context")
	}
	return val, nil
}
