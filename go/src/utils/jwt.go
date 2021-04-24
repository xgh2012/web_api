package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("<44#!@!99>")

type JwtService struct {
	Token string
	jwt.StandardClaims
}

func (j *JwtService) GenerateToken() {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	StandardClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    "gin",
		Audience:  "system",
	}

	j.StandardClaims = StandardClaims

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	j.Token, _ = tokenClaims.SignedString(jwtSecret)
}

//验证token
func (j *JwtService) ParseToken() (*JwtService, error) {
	tokenClaims, err := jwt.ParseWithClaims(j.Token, &JwtService{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JwtService); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
