package model

import "gorm.io/gorm"

type Nodes struct {
	gorm.Model
	SelfId       string
	FatherNodeId string
	NodeType     string //zone或passage
	Content      string
	Likes        int //点赞数量
}
