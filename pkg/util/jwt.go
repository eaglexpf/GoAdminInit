package util

import (
	//	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	//	"github.com/eaglexpf/GoAdminInit/models/auth"
	"github.com/eaglexpf/GoAdminInit/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

type UserData struct {
	UserID int         `json:"user_id"`
	UUID   string      `json:"uuid"`
	Data   interface{} `json:"data"`
}

type Claims struct {
	UserData
	jwt.StandardClaims
}

func GenerateToken(user_id int, uuid string, data interface{}) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		UserData{
			UserID: user_id,
			UUID:   uuid,
			Data:   data,
		},
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "goAdmin",
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
