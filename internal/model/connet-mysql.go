package model

import (
	config "NGB-SE/internal/conf"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//
//迁移
//
func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Nodes{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&LikeAction{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&UserRelation{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&FriendAction{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&InStationMessage{})
	if err != nil {
		return err
	}
	return nil
}

var db *gorm.DB
var err error

func init() {
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", config.Config.DataBase.UserName, config.Config.DataBase.PassWord, config.Config.DataBase.Host, config.Config.DataBase.Port, config.Config.DataBase.DbName, config.Config.DataBase.TimeOut)
	//fmt.Println(dsn)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	err = migrate(db)
	if err != nil {
		panic(err)
	}
}
