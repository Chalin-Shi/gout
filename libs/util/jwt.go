package util

import (
	// "crypto/md5"
	// "encoding/hex"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"gout/libs/setting"
)

type Claims struct {
	ID        int    `json:"id"`
	CLUSTERID string `json:"clusterId"`
	jwt.StandardClaims
}

func GenerateToken(id int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * 24 * time.Hour)

	claims := Claims{
		id,
		"",
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "linktimecloud",
		},
	}

	var jwtSecret = []byte(setting.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string, secret string) (*Claims, error) {
	var jwtSecret = []byte(secret)
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
