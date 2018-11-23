package model

import (
	"log"
	"meeting/bootstrap"
	"meeting/service"
	"time"
)

type Department struct {
	Id     			int
	Department_id   int
	Name            string
	Parent_id       int
	Create_time     int
}

func NewDepartment()*Department  {
	return new(Department)
}

func (this *Department)GetDepartmentList() bool {
	// 部门列表
	result := service.NewWx().GetDepartmentList()
	// 更新部门
	go this.departmentInit(result)
	// 更新会议类型
	go this.meetingTypeInit(result)

	return true;
}

func(this *Department) departmentInit(result service.DepartmentList)  {
	// 删除数据库
	_,err := bootstrap.Rose().Table("department").Force().Delete()
	for _,val := range result.Department{
		data := map[string]interface{}{
			"department_id" : val.Id,
			"name"          : val.Name,
			"parent_id"     : val.ParentId,
			"create_time"   : time.Now().Unix(),
		}
		if err != nil {
			log.Fatal(err)
			break
		}
		bootstrap.Rose().Table("department").Data(data).InsertGetId()
	}
}

func (this *Department)meetingTypeInit(result service.DepartmentList)  {
	// 删除数据库
	bootstrap.Rose().Table("meeting_type").Force().Delete()
	for _ , val := range result.Department{
		data := map[string]interface{}{
			"title"         : val.Name + "会议",
			"create_time"   : time.Now().Unix(),
			"department_id" : val.Id,
		}
		bootstrap.Rose().Table("meeting_type").Data(data).InsertGetId()
	}
}

