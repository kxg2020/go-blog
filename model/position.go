package model

import "meeting/bootstrap"

type Position_Meeting_Map struct {
	Id              int64  `orm:"id"              json:"id"`
	Position_name   string `orm:"position_name"   json:"position_name"`
	Meeting_type_id int64  `orm:"meeting_type_id" json:"meeting_type_id"`
}

func NewPosition()*Position_Meeting_Map  {
	return new(Position_Meeting_Map)
}

func (this *Position_Meeting_Map)PositionList()([]Position_Meeting_Map,error)  {
	var position []Position_Meeting_Map
	err := bootstrap.Rose().Table(&position).Group("position_name").Select()
	if err != nil {
		return position,err
	}
	return position,nil
}