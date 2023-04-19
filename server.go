package main

import (
	"NGB-SE/model"
	"NGB-SE/view"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "The server is running")
	})
	//读取配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	file, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = viper.ReadConfig(file)
	username := viper.GetString("DataBase.username")
	password := viper.GetString("DataBase.password")
	host := viper.GetString("DataBase.host")
	port := viper.GetString("DataBase.port")
	dbName := viper.GetString("DataBase.dbName")
	timeout := viper.GetString("DataBase.timeout")
	if err != nil {
		panic(err)
	}
	model.ConnectDB(username, password, host, port, dbName, timeout) //use model.DB to operate
	view.SetRouters(e)
	e.Logger.Fatal(e.Start(":1323"))

}
