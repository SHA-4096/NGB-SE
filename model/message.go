package model

//
//创建从fromId发送到toId的信息
//
func CreateMessage(fromId, toId, message string) error {
	inStationMessage := new(InStationMessage)
	inStationMessage.FromId = fromId
	inStationMessage.ToId = toId
	inStationMessage.Message = message
	db.Create(&inStationMessage)
	return nil
}

//
//获取发送给Uid的所有信息
//
func GetMessage(Uid string) ([]InStationMessage, error) {
	var inStationMessages []InStationMessage
	db.Find(&inStationMessages, "to_id = ?", Uid)
	return inStationMessages, nil
}
