package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}

var DB *gorm.DB
var err error

func ConnectDB(username, password, host, port, dbName, timeout string) {
	//配置MySQL连接参数

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbName, timeout)
	fmt.Println(dsn)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	err = migrate(DB)
	if err != nil {
		panic(err)
	}
}
