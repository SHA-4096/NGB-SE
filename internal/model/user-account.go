package model

import "fmt"

func QueryUidAndPassword(Uid, PasswordEncoded string) (*User, error) {
	var user User
	err := db.Where("Uid = ? AND Password= ?", Uid, PasswordEncoded).First(&user).Error
	return &user, err
}

//
//Query all users that have email subscription
//
func QueryAllSubscriptors() ([]User, error) {
	var users []User
	err := db.Where("Subscribe= ?", true).First(&users).Error
	return users, err
}

func CreateUser(Receiver *User) error {
	err := db.Create(&Receiver).Error
	return err
}

func SaveUser(Receiver *User) error {
	err := db.Save(Receiver).Error
	return err
}

func QueryUid(Uid string) (*User, error) {
	var user User
	err := db.Where("Uid = ?", Uid).First(&user).Error
	return &user, err
}

func DeleteUid(Uid string) error {
	err = db.Where("Uid = ?", Uid).Delete(&User{}).Error
	return err
}

//
//查询用户的Uid对应的邮箱
//如果为空则返回错误
//
func QueryUidForEmail(Uid string) (string, error) {
	var user User
	err = nil
	err = db.Where("Uid = ?", Uid).First(&user).Error
	if user.Email == "" {
		return "", fmt.Errorf("您似乎没有有效的email")
	}
	return user.Email, err
}
