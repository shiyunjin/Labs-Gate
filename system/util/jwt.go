package util

import (
	"github.com/shiyunjin/SchoolNetwork/system/config"
	"gopkg.in/mgo.v2/bson"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

)

var jwtSecret = []byte(config.Get("jwt.secret").(string))

type Claims struct {
	Id 			bson.ObjectId
	Username 	string
	Hash 		string
	Auth 		string
	jwt.StandardClaims
}

func GenerateToken(id bson.ObjectId, username, hash, auth string, exp int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(exp) * time.Hour)

	claims := Claims{
		id,
		username,
		hash,
		auth,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "syj-schoolnetwork",
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