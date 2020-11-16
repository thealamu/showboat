package main

import (
	"context"
	"sync"
)

const (
	goodTestUser     = "user1369"
	goodTestPassword = "user1369Password"
)

//TestDB is a test implementation of the DB interface
type TestDB struct {
	sync.RWMutex
	data map[UserID]Password
}

func NewTestDB() *TestDB {
	data := make(map[UserID]Password)
	data[goodTestUser] = goodTestPassword
	return &TestDB{data: data}
}

func (t *TestDB) MatchCredentials(ctx context.Context, uid UserID, pwd Password) (bool, error) {
	exists, err := t.HasUser(ctx, uid)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, ErrUnknownUser
	}

	t.RLock()
	defer t.RUnlock()
	return t.data[uid] == pwd, nil
}

func (t *TestDB) HasUser(ctx context.Context, uid UserID) (bool, error) {
	t.RLock()
	defer t.RUnlock()
	_, ok := t.data[uid]
	return ok, nil
}

func (t *TestDB) CreateUser(ctx context.Context, uid UserID, pwd Password) error {
	exists, _ := t.HasUser(nil, uid)
	if exists {
		return ErrUserExists
	}
	t.Lock()
	defer t.Unlock()
	t.data[uid] = pwd
	return nil
}

func (t *TestDB) GetPortfolio(ctx context.Context, uid UserID) (Portfolio, error) {
	if uid == goodTestUser {
		return getTestPortfolio(), nil
	}

	exists, _ := t.HasUser(nil, uid)
	if !exists {
		return Portfolio{}, ErrUnknownUser
	}

	return Portfolio{}, nil
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
