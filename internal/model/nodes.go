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
//保存对节点的修改
//
func SaveNode(Receiver *Nodes) error {
	err := db.Save(&Receiver).Error
	return err
}

//
//删除节点
//
func DeleteNode(NodeId string) error {
	var node Nodes
	err := db.Where("self_id = ?", NodeId).First(&node).Error
	if err != nil {
		return err
	}
	if node.NodeType == "passage" {
		/*节点类型为帖子(passage)的时候直接删除*/
		err = db.Where("self_id = ?", NodeId).Delete(&Nodes{}).Error
	} else {
		/*节点类型为分区(zone)的时候删除其下的所有passage以及节点本身*/
		err = db.Where("father_id = ?", NodeId).Delete(&Nodes{}).Error
		if err != nil {
			return err
		}
		err = db.Where("self_id = ?", NodeId).Delete(&Nodes{}).Error
	}
	return err
}

//
//返回一个随机的节点ID字符串
//
func GetRandomId() string {
	rand.Seed(time.Now().Unix())
	randomId := rand.Int()
	var nodes Nodes
	for {
		err = db.Where("self_id= ?", randomId).First(&nodes).Error
		if err == nil { //若这个id已经存在
			rand.Seed(time.Now().Unix())
			randomId = rand.Int()
		} else {
			break
		}

	}
	return fmt.Sprintf("%d", randomId)

}

//
//查询一个nodeId对应的node名称
//
func GetNodeName(nodeId string) (string, error) {
	var nodes Nodes
	err = db.Where("self_id = ?", nodeId).First(&nodes).Error
	if err != nil {
		return "", fmt.Errorf("这个id不存在")
	}
	return nodes.NodeName, nil
}

//
//查询所有分区
//
func GetAllZones() ([]Nodes, error) {
	var nodes []Nodes
	db.Find(&nodes, "node_type = ?", "zone")
	return nodes, nil

}

//
//获得zoneId对应分区下面的所有文章
//
func GetAllPassageByZones(zoneId string) ([]Nodes, error) {
	var nodes []Nodes
	db.Find(&nodes, "node_type = ? AND father_node_id = ?", "passage", zoneId)
	return nodes, nil
}

//
//根据nodeId返回一个node结构体
//
func GetSingleNode(nodeId string) (Nodes, error) {
	var node Nodes
	err := db.Where("self_id = ?", nodeId).First(&node).Error
	return node, err
}
