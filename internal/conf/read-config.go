package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type configStruct struct {
	DataBase    dataBaseStruct    `yaml:"DataBase"`
	JWTConfig   jWTConfigStruct   `yaml:"JWTConfig"`
	LogConfig   logConfigStruct   `yaml:"LogConfig"`
	EmailConfig emailConfigStruct `yaml:"EmailConfig"`
}

type dataBaseStruct struct {
	UserName      string `yaml:"username"`
	PassWord      string `yaml:"password"`
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	DbName        string `yaml:"dbName"`
	TimeOut       string `yaml:"timeout"`
	RedisAddr     string `yaml:"redisAddr"`
	RedisPassword string `yaml:"redisPassword"`
	RedisDB       int    `yaml:"redisDB"`
}

type jWTConfigStruct struct {
	RefreshTokenKey string `yaml:"refreshTokenKey"`
}

type logConfigStruct struct {
	LogPath    string `yaml:"logPath"`
	RotateTime int    `yaml:"rotateTime"`
	MaxAge     int    `yaml:"maxAge"`
}

type emailConfigStruct struct {
	EmailAddress      string `yaml:"emailAddress"`
	SmtpServer        string `yaml:"smtpServer"`
	SmtpPort          int    `yaml:"smtpPort"`
	Name              string `yaml:"name"`
	Password          string `yaml:"password"`
	ExpirationSeconds int    `yaml:"expirationSeconds"`
	SubscriptionHour  int    `yaml:"subscriptionHour"`
}

var Config configStruct

func init() {
	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic("Error when loading config")
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic("Error when unmarshaling")
	}
	fmt.Println("Config loaded")
}
