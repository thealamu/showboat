package main

import "context"

type DB interface {
	GetPortfolio(context.Context, UserID) (Portfolio, error)
}
