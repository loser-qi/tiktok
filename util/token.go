package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claim struct {
	Id int64
	jwt.StandardClaims
}

var secret = []byte("g_c_s_j")
var expireTime = time.Now().Add(30 * time.Minute)

func EncodeToken(id int64) string {
	claim := &Claim{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
		},
	}
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(secret)
	if err != nil {
		panic("token生成错误, err=" + err.Error())
	}
	return tokenStr
}

func DecodeToken(tokenStr string) (int64, error) {
	if tokenStr == "" {
		return 0, errors.New("token is invalid")
	}
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("token is invalid")
	}
	return claim.Id, nil
}
