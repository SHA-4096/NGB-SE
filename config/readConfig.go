package config

import (
	"os"

	"github.com/spf13/viper"
)

type dataBaseStruct struct {
	UserName string
	PassWord string
	Host     string
	Port     string
	DbName   string
	TimeOut  string
}

type jWTConfigStruct struct {
	RefreshTokenKey string
}

var (
	DataBase  *dataBaseStruct
	JwtConfig *jWTConfigStruct
)

func init() {
	DataBase = new(dataBaseStruct)
	JwtConfig = new(jWTConfigStruct)
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
	DataBase.UserName = viper.GetString("DataBase.username")
	DataBase.PassWord = viper.GetString("DataBase.password")
	DataBase.Host = viper.GetString("DataBase.host")
	DataBase.Port = viper.GetString("DataBase.port")
	DataBase.DbName = viper.GetString("DataBase.dbName")
	DataBase.TimeOut = viper.GetString("DataBase.timeout")
	JwtConfig.RefreshTokenKey = viper.GetString("JwtConfig.refreshTokenKey")
	if err != nil {
		panic(err)
	}
}
