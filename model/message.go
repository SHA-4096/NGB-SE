package model

func CreateMessage(fromId, toId, message string) {
	inStationMessage := new(InStationMessage)
	inStationMessage.FromId = fromId
	inStationMessage.ToId = toId
	inStationMessage.Message = message
}
