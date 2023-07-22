package config

import (
	"os"

	"github.com/spf13/viper"
)

type dataBaseStruct struct {
	UserName      string
	PassWord      string
	Host          string
	Port          string
	DbName        string
	TimeOut       string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

type jWTConfigStruct struct {
	RefreshTokenKey string
}

type logConfigStruct struct {
	LogPath    string
	RotateTime int
	MaxAge     int
}

type emailConfigStruct struct {
	EmailAddress string
	SmtpServer   string
	SmtpPort     int
	Name         string
	Password     string
}

var (
	DataBase    *dataBaseStruct
	JwtConfig   *jWTConfigStruct
	LogConfig   *logConfigStruct
	EmailConfig *emailConfigStruct
)

func init() {
	DataBase = new(dataBaseStruct)
	JwtConfig = new(jWTConfigStruct)
	LogConfig = new(logConfigStruct)
	EmailConfig = new(emailConfigStruct)
	//读取配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	file, err := os.Open("config/config.yaml")
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
	DataBase.RedisAddr = viper.GetString("DataBase.redisAddr")
	DataBase.RedisPassword = viper.GetString("DataBase.redisPassword")
	DataBase.RedisDB = viper.GetInt("DataBase.redisDB")
	JwtConfig.RefreshTokenKey = viper.GetString("JwtConfig.refreshTokenKey")
	LogConfig.LogPath = viper.GetString("LogConfig.logPath")
	LogConfig.RotateTime = viper.GetInt("LogConfig.rotateTime")
	LogConfig.MaxAge = viper.GetInt("LogConfig.maxAge")
	EmailConfig.EmailAddress = viper.GetString("EmailConfig.emailAddress")
	EmailConfig.SmtpPort = viper.GetInt("EmailConfig.smtpPort")
	EmailConfig.SmtpServer = viper.GetString("EmailConfig.smtpServer")
	EmailConfig.Name = viper.GetString("EmailConfig.name")
	EmailConfig.Password = viper.GetString("EmailConfig.password")
	if err != nil {
		panic(err)
	}
}
