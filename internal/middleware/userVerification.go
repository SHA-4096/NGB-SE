package middleware

import (
	"NGB-SE/internal/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func VerifyUser(Uid, tokenRaw string, isRefresh bool) (string, error) {
	/*内部函数，用来验证用户,需要:Uid的路径参数以及token,会返回一个token和error*/
	user, err := model.QueryUid(Uid)
	if err != nil {
		return "", fmt.Errorf("您不是已注册用户")
	}
	token := (strings.Split(tokenRaw, " "))
	if len(token) < 2 {
		return "", fmt.Errorf("你没有在请求头携带token")
	}
	//fmt.Println("TOKEN IS:", token[1])
	var key string
	if isRefresh {
		key = refreshTokenKey
	} else {
		key = user.JwtKey
	}

	claims, err := DecodeJwt(token[1], key)
	if err != nil {
		return "", err
	}
	if claims["Uid"] != Uid {
		return "", fmt.Errorf("你的token无效")
	}
	return token[1], nil
}

func VerifyAdmin(AdminId, tokenRaw string) error {
	token := (strings.Split(tokenRaw, " "))
	if len(token) < 2 {
		return fmt.Errorf("你没有在请求头携带token")
	}
	//fmt.Println("TOKEN IS:", token[1])
	//检查管理员身份
	user, err := model.QueryUid(AdminId)
	if err != nil {
		return err
	}
	key := user.JwtKey
	claims, err := DecodeJwt(token[1], key)
	if err != nil {
		return err
	}
	if claims["Uid"] != AdminId {
		return fmt.Errorf("你的token似乎和你的身份不符")
	}
	if user.IsAdmin != "True" {
		return fmt.Errorf("你不是管理员")
	}
	return nil
}

func EncodeMethod(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
