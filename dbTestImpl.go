package main

import (
	"context"
	"errors"
)

//TestDB is a test implementation of the DB interface
type TestDB struct{}

const goodTestUser = "user1369"

func (t TestDB) GetPortfolio(ctx context.Context, uid UserID) (Portfolio, error) {
	if uid != goodTestUser {
		return Portfolio{}, errors.New("user does not exist")
	}

	return getTestPortfolio(), nil
}

func getTestPortfolio() Portfolio {
	return Portfolio{
		[]PortfolioItem{
			{
				"https://via.placeholder.com/150",
				"FooBar Project",
				"Some foo bar project",
				nil,
				"Some nice content",
			},
			{
				"https://via.placeholder.com/300",
				"Big Project",
				"A big project",
				[]Url{"https://via.placeholder.com/300", "https://via.placeholder.com/150"},
				"Big content",
			},
			{
				"https://via.placeholder.com/500",
				"Bigger Project",
				"A bigger project",
				[]Url{"https://via.placeholder.com/500", "https://via.placeholder.com/300", "https://via.placeholder.com/500"},
				"Bigger content",
			},
		},
	}
}
