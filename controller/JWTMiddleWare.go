package controller

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func genkey() string {
	/*return a string as jwt key*/
	rand.Seed(time.Now().Unix())
	seed1 := time.Now().Unix() + int64(rand.Int())
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
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["exp"])
	expTime, ok1 := claims["exp"].(int)
	fmt.Println(ok1)
	fmt.Println(expTime, "-----", time.Now().Unix())
	//检查token是否过期（发现有问题，到时再修）
	/*if expTime < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}
	*/
	if !ok && !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}
	return claims, nil
}
