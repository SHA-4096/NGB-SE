package controller

import (
	"NGB-SE/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	/*POST Uid;Name;Password*/
	uid := c.FormValue("Uid")
	name := c.FormValue("Name")
	password := c.FormValue("password")
	token, key, err := GetJwt(uid)
	if err != nil {
		panic(err)
	}

	user := model.User{
		Name:     name,
		Uid:      uid,
		Password: password,
		Jwt_key:  key,
	}
	model.DB.Create(&user)
	return c.String(http.StatusOK, "JWT:"+token)
}
