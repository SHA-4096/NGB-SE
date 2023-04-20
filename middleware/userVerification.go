package middleware

import (
	"NGB-SE/model"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

func VerifyUser(c echo.Context, isRefresh bool) (string, error) {
	/*内部函数，用来验证用户,需要:Uid的路径参数以及token,会返回一个token和error*/
	Uid := c.Param("Uid")
	user := new(model.User)
	model.DB.Where("Uid = ?", Uid).First(&user)
	if user.Uid == "" {
		return "", fmt.Errorf("您不是已注册用户")
	}
	tokenRaw := c.Request().Header.Get("Authorization")
	token := (strings.Split(tokenRaw, " "))
	if len(token) < 2 {
		return "", fmt.Errorf("你没有在请求头携带token")
	}
	fmt.Println("TOKEN IS:", token[1])
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

func VerifyAdmin(c echo.Context) error {
	AdminId := c.Param("AdminId")
	user := new(model.User)
	tokenRaw := c.Request().Header.Get("Authorization")
	token := (strings.Split(tokenRaw, " "))
	if len(token) < 2 {
		return fmt.Errorf("你没有在请求头携带token")
	}
	fmt.Println("TOKEN IS:", token[1])
	//检查管理员身份
	model.DB.Where("Uid = ?", AdminId).First(&user)
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
