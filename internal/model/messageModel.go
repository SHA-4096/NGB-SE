package model

//这个model负责站内信息的存储

type InStationMessage struct {
	FromId  string
	ToId    string
	Message string
}
