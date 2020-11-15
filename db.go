package main

import "context"

type DB interface {
	GetPortfolio(context.Context, UserID) (Portfolio, error)
	HasUser(context.Context, UserID) (bool, error)
	CreateUser(context.Context, UserID, Password) error
}
