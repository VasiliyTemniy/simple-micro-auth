package services

import (
	"crypto/rsa"
	"fmt"
	"math/rand"
	"simple-micro-auth/src/cert"
	m "simple-micro-auth/src/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenHandler interface {
	CreateToken(id int64, timeNow int64, expiresAt int64, noise int64, key *rsa.PrivateKey) (string, error)
	VerifyToken(token string, timeNow int64, key interface{}) (int64, error)
	RefreshToken(token string, tokenTtl time.Duration) *m.AuthResponse
}

type tokenHandlerImpl struct{}

func (handler *tokenHandlerImpl) CreateToken(id int64, timeNow int64, expiresAt int64, noise int64, key *rsa.PrivateKey) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":   id,
		"rand": noise,
		"exp":  expiresAt,
		"iat":  timeNow,
		"nbf":  timeNow,
		"iss":  "simple-micro-auth",
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (handler *tokenHandlerImpl) VerifyToken(token string, timeNow int64, key interface{}) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		err = fmt.Errorf("invalid token or signature")
		return 0, err
	}

	if !parsedToken.Valid {
		err = fmt.Errorf("invalid token")
		return 0, err
	}

	if parsedToken.Method != jwt.SigningMethodRS256 {
		err = fmt.Errorf("unexpected signing method")
		return 0, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("claims malformed")
		return 0, err
	}

	expiresAt, ok := claims["exp"].(float64)
	if !ok {
		err = fmt.Errorf("exp claim malformed")
		return 0, err
	}

	if int64(expiresAt) < timeNow {
		err = fmt.Errorf("token expired")
		return 0, err
	}

	id, ok := claims["id"].(float64)
	if !ok {
		err = fmt.Errorf("id claim malformed")
		return 0, err
	}

	return int64(id), nil
}

func (handler *tokenHandlerImpl) RefreshToken(token string, tokenTtl time.Duration) *m.AuthResponse {

	id, err := handler.VerifyToken(token, time.Now().Unix(), cert.PublicKey)
	if err != nil {
		return &m.AuthResponse{
			Id:    id,
			Token: token,
			Error: "TokenError: " + err.Error(),
		}
	}

	tokenString, err := handler.CreateToken(id, time.Now().Unix(), time.Now().Add(tokenTtl).Unix(), rand.Int63(), cert.PrivateKey)
	if err != nil {
		return &m.AuthResponse{
			Id:    id,
			Token: tokenString,
			Error: "TokenError: " + err.Error(),
		}
	}

	return &m.AuthResponse{
		Id:    id,
		Token: tokenString,
		Error: "",
	}
}

func NewTokenHandler() ITokenHandler {
	return &tokenHandlerImpl{}
}
