package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)


type CaptorInfo struct {
	Id int `orm:"pk"`
	Url string `orm:"size(255)"`
	Desc string `orm:"size(255)"`
	Keyword string `orm:"size(255)"`
	Createat time.Time `orm:"type(datetime)"`
	Updateat time.Time `orm:"type(datetime)"`
}

func (m *CaptorInfo) TableName() string {
	return TableName("captorinfo")
}

func (m *CaptorInfo) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err != nil {
		return err
	}
	return nil
}


func (m *CaptorInfo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields ...);err != nil {
		return err
	}
	return nil
}

func (m *CaptorInfo) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields ...);err != nil {
		return err
	}
	return nil
}

func (m *CaptorInfo) Delete() error {
	if _,err := orm.NewOrm().Delete(m);err != nil {
		return err
	}
	return nil
}

func (m *CaptorInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *CaptorInfo) FindAll(container interface{}) bool{
	t,err := m.Query().All(container)
	if err == nil && t > 0 {
		return true
	}
	return false
}

func (m *CaptorInfo) ToJsonString() string {
	data,_ := json.Marshal(m)
	return string(data)
}