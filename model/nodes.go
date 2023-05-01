package model

import (
	"fmt"
	"math/rand"
	"time"
)

//
//创建节点
//
func CreateNode(Receiver *Nodes) error {
	err := db.Create(&Receiver).Error
	return err
}

//
//删除节点
//
func DeleteNode(NodeId string) error {
	var node Nodes
	err := db.Where("SelfId = ?", NodeId).First(&nodes).Error
	if err != nil {
		return err
	}
	if node.NodeType == "passage" {
		/*节点类型为帖子(passage)的时候直接删除*/
		err = db.Where("SelfId = ?", NodeId).Delete(&User{}).Error
	} else {
		/*节点类型为分区(zone)的时候删除其下的所有passage以及节点本身*/
		err = db.Where("FatherId = ?", NodeId).Delete(&User{}).Error
		if err != nil {
			return err
		}
		err = db.Where("SelfId = ?", NodeId).Delete(&User{}).Error
	}
	return err
}

//
//返回一个随机的节点ID字符串
//
func GetRandomId() string {
	rand.Seed(time.Now().Unix())
	randomID := rand.Int()
	var nodes Nodes
	for {
		err = db.Where("SelfId = ?", randomID).First(&nodes).Error
		if err != nil {
			rand.Seed(time.Now().Unix())
			randomID = rand.Int()
		} else {
			break
		}

	}
	return fmt.Sprintf("%d", randomID)

}
