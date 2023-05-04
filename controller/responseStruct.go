package controller

type CreateNodeInData struct {
	ZoneName string
	ZoneId   string
	Content  string
}

type ZoneNameStruct struct {
	ZoneId   string
	ZoneName string
}

type PassageStruct struct {
	PassageId      string
	PassageContent string
	BelongZoneName string
	BelongZoneId   string
}

type MsgStruct struct {
	Message string
}
