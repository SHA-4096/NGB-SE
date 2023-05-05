package model

import "fmt"

//
//查询用户点赞的所有文章节点记录
//
func GetAllLikes(Uid string) ([]LikeAction, error) {
	var actions []LikeAction
	db.Find(&actions, "uid = ?", Uid)
	return actions, nil
}

//
//创建点赞记录
//
func CreateLike(Uid string, PassageId string) error {
	var action LikeAction
	err := db.Where("uid = ? AND passage_id = ?", Uid, PassageId).First(&action).Error
	if err != nil {
		//找不到记录
		action.Uid = Uid
		action.PassageId = PassageId
		db.Create(&action)
		return nil
	}
	return fmt.Errorf("这条点赞记录已经存在")
}

func DeleteLike(Uid string, PassageId string) error {
	err := db.Where("uid = ? AND passage_id = ?", Uid, PassageId).Delete(&LikeAction{}).Error
	if err != nil {
		return err
	}
	return nil
}