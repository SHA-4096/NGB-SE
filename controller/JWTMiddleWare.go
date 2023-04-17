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
