package model

type UserRelation struct {
	RelationType string //分为follow和friend
	Uid          string //主体的Id
	TargetId     string //客体的Id
}
