package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("ssshhhhfkjk")

type Claims struct {
	Id      string `json:"id"`
	Account string `json:"account"`
	OpenId  string `json:"openId"`
	jwt.StandardClaims
}

func GenToken(id string, account string, openId string) (r string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		id,
		account,
		openId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "yznu-cn",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func GetNewToken(token string) (r string, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok {
			if claims.ExpiresAt < (time.Now().Unix() - 60) {
				err = errors.New("token 已经过期")
				return
			}
			return GenToken(claims.Id, claims.Account, claims.OpenId)
		}
	}
	return "", err
}
