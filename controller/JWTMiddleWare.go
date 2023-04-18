package controller

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func genkey() string {
	/*return a string as jwt key*/
	seed1 := rand.Intn(10000000000000)
	seed2 := "AVerySecureKey:)" //先这样用着吧（）
	return fmt.Sprintf("%d%s", seed1, seed2)
}

const expHours = 1

func GetJwt(Uid string) (string, string, error) {
	/*1st:token 2nd:key*/
	expTime := time.Now().Add(time.Hour * expHours).Unix()
	claims := jwt.MapClaims{}
	claims["Uid"] = Uid
	claims["exp"] = expTime
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := genkey()
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}
	return signedToken, key, nil

}

func DecodeJwt(tokenString, key string) (jwt.MapClaims, error) {
	/*return the claim of a token*/
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(key), nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}
	return claims, nil
}
