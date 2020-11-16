package main

import (
	"context"
	"errors"
	"fmt"
	"regexp"
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

	isValid, err := regexp.MatchString(fmt.Sprintf("^%s$", UserIDFormat), userid)
	if err != nil {
		return "", err
	}
	if !isValid {
		return "", errors.New("userid is invalid")
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

	return token.SignedString([]byte(s.hmacSecret))
}

var ErrBadClaimType = errors.New("bad claim type stored in token")

//claimFromToken returns the claim stored in a token
func (s *Server) claimFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.hmacSecret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("error parsing token")
	}

	claim, ok := claims["uid"].(string)
	if !ok {
		return "", ErrBadClaimType
	}

	return claim, nil
}
