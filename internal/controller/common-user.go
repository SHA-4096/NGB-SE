package controller

import (
	"NGB-SE/internal/controller/param"
	"NGB-SE/internal/middleware"
	"NGB-SE/internal/model"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

type LogOutInData struct {
	Uid string
}

func Register(c echo.Context) error {
	/*POST Uid;Name;Password*/
	//检查id是否被使用
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	_, err := model.QueryUid(user.Uid)
	if err == nil {
		outData := map[string]interface{}{
			"message": "这个用户id已经被占用了哦",
		}
		return c.JSON(http.StatusBadRequest, outData)

	}
	user.IsAdmin = "False"
	user.Password = middleware.EncodeMethod(user.Password)
	model.CreateUser(user)

	data := map[string]interface{}{
		"message": "注册成功",
	}
	return c.JSON(http.StatusCreated, data)
}

func Login(c echo.Context) error {
	/*POST Uid;Password*/
	inData := new(model.User)
	c.Bind(inData)
	user, err := model.QueryUidAndPassword(inData.Uid, middleware.EncodeMethod(inData.Password))
	if err != nil {
		outData := map[string]interface{}{
			"message": "帐号不存在或密码错误",
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	jwtToken, key, err := middleware.GetJwt(user.Uid)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	user.JwtKey = key
	model.SaveUser(user)
	refreshToken, err := middleware.GetRefreshJwt(user.Uid)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	outData := map[string]interface{}{
		"jwtToken":     jwtToken,
		"refreshToken": refreshToken,
		"message":      fmt.Sprintf("welcome:%s", user.Name),
	}
	return c.JSON(http.StatusOK, outData)

}

func LogOut(c echo.Context) error {
	/*GET src = /user/:Uid/logout with token*/
	tokenRaw := c.Request().Header.Get("Authorization")
	_, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, false)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	//从数据库里面取出用户
	user, err := model.QueryUid(c.Param("Uid"))
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	//注销用户
	user.JwtKey = ""
	model.SaveUser(user)
	outData := map[string]interface{}{
		"message": fmt.Sprintf("用户%s已经注销", c.Param("Uid")),
	}
	return c.JSON(http.StatusOK, outData)
}

func DeleteUser(c echo.Context) error {
	/*GET src = /user/delete/:Uid with token*/
	tokenRaw := c.Request().Header.Get("Authorization")
	_, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, false)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	err = model.DeleteUid(c.Param("Uid"))
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	} else {
		outData := map[string]interface{}{
			"message": fmt.Sprintf("用户%s已经删除", c.Param("Uid")),
		}
		return c.JSON(http.StatusOK, outData)
	}

}

func ModifyUser(c echo.Context) error {
	/*POST src = /user/modify/:Uid with json containing key&value*/
	inData := new(AdminModifyUserINData)
	//验证用户
	tokenRaw := c.Request().Header.Get("Authorization")
	_, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, false)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	//获取数据
	err = c.Bind(inData)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	user, err := model.QueryUid(c.Param("Uid"))
	//查询出错时
	if user.Uid == "" {
		outData := map[string]interface{}{
			"message": fmt.Sprintf("数据库出了问题耶，错误信息：%s", err),
		}
		return c.JSON(http.StatusBadRequest, outData)
	}
	//找到用户时
	refUser := reflect.ValueOf(user).Elem()
	fieldValue := refUser.FieldByName(inData.Key)
	if fieldValue.IsValid() {
		fieldValue.SetString(inData.Value)
		model.SaveUser(user)
		outData := map[string]interface{}{
			"message": fmt.Sprintf("用户%s的%s值被修改为%s", c.Param("Uid"), inData.Key, inData.Value),
		}
		return c.JSON(http.StatusOK, outData)
	} else {
		outData := map[string]interface{}{
			"message": "要修改的键值不存在",
		}
		return c.JSON(http.StatusBadRequest, outData)
	}

}

func RenewWithRefreshToken(c echo.Context) error {
	/*GET方法,更新jwtToken,src = /user/:Uid/refreshtoken,请求头携带refreshToken*/
	//检查用户
	tokenRaw := c.Request().Header.Get("Authorization")
	refreshToken, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, true)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}

	user, err := model.QueryUid(c.Param("Uid"))
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	//分发新的token
	jwtToken, key, err := middleware.GetJwt(user.Uid)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	user.JwtKey = key
	model.SaveUser(user)
	outData := map[string]interface{}{
		"jwtToken":     jwtToken,
		"refreshToken": refreshToken,
		"message":      fmt.Sprintf("welcome:%s", user.Name),
	}
	return c.JSON(http.StatusOK, outData)

}

func GetLoginCode(c echo.Context) error {
	inData := new(model.User)
	c.Bind(inData)
	email, err := model.QueryUidForEmail(inData.Uid)
	if err == nil {
		//找到邮箱
		err = middleware.SendVerificationEmail(email)
	}
	if err != nil {
		//找不到对应邮箱或信息发送失败
		outData := map[string]interface{}{
			"status": "fail",
			"data":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}

	outData := map[string]interface{}{
		"status": "success",
		"data":   "",
	}
	return c.JSON(http.StatusOK, outData)
}

func SendLoginCode(c echo.Context) error {
	inData := new(param.RequestCodeLogin)
	c.Bind(inData)
	email, err := model.QueryUidForEmail(inData.Uid)
	if err != nil {
		//找不到对应邮箱
		outData := map[string]interface{}{
			"status": "fail",
			"data":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	err = middleware.VerifyUserByEmail(email, inData.Code)
	if err != nil {
		//验证失败
		outData := map[string]interface{}{
			"status": "fail",
			"data":   err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	//获取用户结构体
	user, err := model.QueryUid(inData.Uid)
	if err != nil {
		return err
	}
	jwtToken, key, err := middleware.GetJwt(user.Uid)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	user.JwtKey = key
	model.SaveUser(user)
	refreshToken, err := middleware.GetRefreshJwt(user.Uid)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	outData := map[string]interface{}{
		"jwtToken":     jwtToken,
		"refreshToken": refreshToken,
		"message":      fmt.Sprintf("welcome:%s", user.Name),
	}
	return c.JSON(http.StatusOK, outData)
}
