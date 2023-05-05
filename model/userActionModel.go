package model

type LikeAction struct {
	Uid       string
	PassageId string
}

type FriendAction struct {
	//加好友的请求记录，临时存储，同意或者拒绝之后删除
	Uid      string //主体Uid
	TargetId string //客体Uid
}
