package middleware

import (
	config "NGB-SE/internal/conf"
	"NGB-SE/internal/model"
	"NGB-SE/internal/util"
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

//
//邮箱认证，发送验证码部分
//
func SendVerificationEmail(Email string) error {
	_, err := model.GetKeyValue(Email)
	if err == nil {
		return fmt.Errorf("请求过于频繁，请一段时间后再试")
	}

	//生成邮箱验证码
	code := util.GetRamdomCode(6)
	//设置键值对
	err = model.SetKeyValuePair(Email, code)
	if err != nil {
		return err
	}
	//设置过期时间
	err = model.SetExpiration(Email, config.Config.EmailConfig.ExpirationSeconds)
	if err != nil {
		return err
	}
	//发送邮件
	content := fmt.Sprintf("您正在尝试使用邮箱登录到NGB-SE<br>您的验证码为：%s <br><br> 如果不是您本人操作，请忽略此邮件", code)
	err = util.SendEmail(Email, "Verification-code", "text/html", content)
	return err
}

//
//邮箱认证，检验验证码部分
//
func VerifyUserByEmail(Email string, code string) error {
	val, err := model.GetKeyValue(Email)
	if err != nil {
		return err
	}
	if val != code {
		//验证失败
		return fmt.Errorf("验证码不正确或已过期，请重新获取")
	}
	//验证成功
	return nil
}
