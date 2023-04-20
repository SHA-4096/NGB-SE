package controller

import (
	"NGB-SE/config"
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
const refreshExpHours = 100

var refreshTokenKey string

func init() {
	refreshTokenKey = config.JwtConfig.RefreshTokenKey
}

func GetJwt(Uid string) (string, string, error) {
	/*1st:token 2nd:key*/
	key := genkey()
	expTime := time.Now().Add(time.Hour * expHours).Unix()
	//生成jwtToken
	jwtClaims := jwt.MapClaims{}
	jwtClaims["Uid"] = Uid
	jwtClaims["exp"] = expTime
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	jwtToken, err := tokenJwt.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}

	return jwtToken, key, nil

}

func GetRefreshJwt(Uid string) (string, error) {
	//生成refreshToken
	fmt.Println("RefreshTokenKeyIs", refreshTokenKey)
	refreshExpTime := time.Now().Add(time.Hour * refreshExpHours).Unix()
	refreshClaims := jwt.MapClaims{}
	refreshClaims["Uid"] = Uid
	refreshClaims["exp"] = refreshExpTime
	refreshClaims["Randpayload"] = genkey() //用genkey弄一个新的随机数
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := tokenRefresh.SignedString([]byte(refreshTokenKey))
	if err != nil {
		panic(err)
	}
	return refreshToken, nil
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
	//fmt.Println(claims["exp"])
	expTime, _ := claims["exp"].(float64)
	//fmt.Println(ok1)
	//fmt.Println(int64(expTime), "-----", time.Now().Unix())
	//检查token是否过期（发现有问题，到时再修）
	if int64(expTime) < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}

	if !ok && !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}
	return claims, nil
}
