package model

import "fmt"

const (
	//二进制状态码
	StateFriend = 1 //1
	StateFollow = 2 //10
)

//
//查询用户点赞的所有文章节点记录
//
func GetAllLikes(Uid string) ([]LikeAction, error) {
	var actions []LikeAction
	db.Find(&actions, "uid = ?", Uid)
	return actions, nil
}

//
//查询用户关注列表
//
func GetAllFollows(Uid string) ([]UserRelation, error) {
	var relations []UserRelation
	db.Find(&relations, "uid = ? AND relation_type = ?", Uid, "follow")
	return relations, nil
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

//
//创建好友关系
//
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

//
//创建关注关系
//
func CreateFollowRelation(Uid, FollowId string) error {
	userRelation := UserRelation{
		RelationType: "follow",
		Uid:          Uid,
		TargetId:     FollowId,
	}
	err := db.Create(&userRelation).Error
	return err

}

//
//
//
func DeleteFollowRelation(Uid, FollowId string) error {
	err := db.Where("uid = ? AND target_id = ? AND relation_type = ?", Uid, FollowId, "follow").Delete(&UserRelation{}).Error
	return err
}

//
//获取用户的关系
//返回一个二进制状态码
//好友：1 关注：10
//
func GetUserRelation(Uid, TargetId string) (int, error) {
	var userRelation UserRelation
	returnState := 0
	err := db.Where("uid = ? AND target_id = ? AND relation_type = ?", Uid, TargetId, "friend").First(&userRelation).Error
	if err == nil {
		returnState += StateFriend
	}
	err = db.Where("uid = ? AND target_id = ? AND relation_type = ?", Uid, TargetId, "follow").First(&userRelation).Error
	if err == nil {
		returnState += StateFollow
	}
	return returnState, nil

}
