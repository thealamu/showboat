package main

import (
	"context"
	"errors"
)

var (
	ErrUserExists  = errors.New("user exists")
	ErrUnknownUser = errors.New("user does not exist")
)

type DB interface {
	GetPortfolio(context.Context, UserID) (Portfolio, error)
	HasUser(context.Context, UserID) (bool, error)
	CreateUser(context.Context, UserID, Password) error
}
