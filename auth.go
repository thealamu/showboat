package main

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//signupUser creates a new user with the credentials and returns a token
func (s *Server) signupUser(ctx context.Context, userid, password string) (string, error) {
	userid = strings.TrimSpace(userid)
	password = strings.TrimSpace(password)

	if userid == "" || password == "" {
		return "", errors.New("userid and password are required")
	}

	userExists, err := s.db.HasUser(ctx, userid)
	if err != nil {
		s.logger.Error(err)
		return "", err
	}

	if userExists {
		return "", errors.New("userid exists")
	}

	s.logger.Infof("Creating user '%s'", userid)
	err = s.db.CreateUser(ctx, userid, password)
	if err != nil {
		s.logger.Error(err)
		return "", err
	}

	return s.createToken(userid)
}

//createToken returns a jwt token claiming a user ID
func (s *Server) createToken(userid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "showboat",
		"uid": userid,
	})

	hmacSecret, ok := os.LookupEnv("HMACSECRET")
	if !ok {
		s.logger.Fatal("hmacsecret not found in env")
	}

	return token.SignedString([]byte(hmacSecret))
}
