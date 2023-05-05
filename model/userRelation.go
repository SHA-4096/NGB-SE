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

func CreateFirendRequest(Uid, FriendId string) error {
	friendAction := FriendAction{
		Uid:      Uid,
		TargetId: FriendId,
	}
	err := db.Where("uid = ? AND target_id = ?", Uid, FriendId).First(&friendAction).Error
	if err == nil {
		return fmt.Errorf("你已经发出了请求")
	}
	err = db.Where("uid = ? AND target_id = ?", FriendId, Uid).First(&friendAction).Error
	if err == nil {
		return fmt.Errorf("对方已经发出了请求")
	}
	db.Create(&friendAction)

	return nil
}

//
//检查Uid是否向FriendId发送了请求，如果有的话就返回真
//
func CheckFriendRequest(Uid, FriendId string) error {
	var friendAction FriendAction
	err := db.Where("uid = ? AND target_id = ?", Uid, FriendId).First(&friendAction).Error
	return err

}

func CreateFriendRelation(Uid, FriendId string) error {
	friendRelation := UserRelation{
		Uid:          Uid,
		TargetId:     FriendId,
		RelationType: "friend",
	}
	db.Create(&friendRelation)
	friendRelation.Uid = FriendId
	friendRelation.TargetId = Uid
	db.Create(&friendRelation)
	return nil
}
