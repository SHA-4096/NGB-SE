package model

func QueryUidAndPassword(Uid, PasswordEncoded string) (*User, error) {
	var user User
	err := db.Where("Uid = ? AND Password= ?", Uid, PasswordEncoded).First(&user).Error
	return &user, err
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
