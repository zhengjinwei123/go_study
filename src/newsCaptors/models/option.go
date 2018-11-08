package models

import (
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type Option struct {
	Id int `orm:"pk"`
	Name string `orm:"size(30)"`
	Value string
}

func (m *Option) TableName() string {
	return TableName("option")
}

func (m *Option) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err != nil {
		return err
	}
	return nil
}


func (m *Option) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...);err != nil {
		return err
	}
	return nil
}

func (m *Option) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields ...);err != nil {
		return err
	}
	return nil
}

func (m *Option) Delete() error {
	if _,err := orm.NewOrm().Delete(m);err != nil {
		return err
	}
	return nil
}

func (m *Option) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Option) FindAll(container interface{}) bool{
	t,err := m.Query().All(container)
	if err == nil && t > 0 {
		return true
	}
	return false
}

func (m *Option) ToJsonString() string {
	data,_ := json.Marshal(m)
	return string(data)
}

