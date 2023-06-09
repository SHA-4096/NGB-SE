package model

import "gorm.io/gorm"

type Nodes struct {
	gorm.Model
	SelfId       string
	FatherNodeId string
	AuthorId     string //文章或分区的作者
	NodeType     string //zone或passage
	NodeName     string //节点名称
	Content      string
	Likes        int //点赞数量
}
