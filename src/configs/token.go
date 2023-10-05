package configs

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenHandler interface {
	CreateToken(id int64) (string, error)
	VerifyToken(token string) (int64, error)
	RefreshToken(token string) *AuthResponse
}

type tokenHandlerImpl struct{}

func (handler *tokenHandlerImpl) CreateToken(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         id,
		"rand":       rand.Intn(100000),
		"expires_at": time.Now().Add(time.Hour * time.Duration(EnvJWTExpirationHours)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(EnvJWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken implements dbHandler.
func (handler *tokenHandlerImpl) VerifyToken(token string) (int64, error) {

	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(EnvJWTSecret), nil
	})

	if err != nil {
		return 0, err
	} else if !tokenClaims.Valid {
		err = fmt.Errorf("invalid token")
		return 0, err
	}

	expiresAt, ok := tokenClaims.Claims.(jwt.MapClaims)["expires_at"].(int64)
	if !ok {
		err = fmt.Errorf("expires_at claim malformed")
		return 0, err
	} else if time.Unix(expiresAt, 0).Before(time.Now()) {
		err = fmt.Errorf("token expired")
		return 0, err
	}

	return tokenClaims.Claims.(jwt.MapClaims)["id"].(int64), nil
}

// RefreshToken implements dbHandler.
func (handler *tokenHandlerImpl) RefreshToken(token string) *AuthResponse {

	id, err := handler.VerifyToken(token)
	if err != nil {
		return &AuthResponse{
			Id:    id,
			Token: token,
			Error: "TokenError: " + err.Error(),
		}
	}

	tokenString, err := handler.CreateToken(id)
	if err != nil {
		return &AuthResponse{
			Id:    id,
			Token: tokenString,
			Error: "TokenError: " + err.Error(),
		}
	}

	return &AuthResponse{
		Id:    id,
		Token: tokenString,
		Error: "",
	}
}

func NewTokenHandler() ITokenHandler {
	return &tokenHandlerImpl{}
}
