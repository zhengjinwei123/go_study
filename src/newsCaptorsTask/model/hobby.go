package model

import (
	"time"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	_ "fmt"
)


type Hobby struct {
	Id int `orm:"pk"`
	Hobby string `orm:"size(255)"`
	Type int
	Createat time.Time `orm:"type(datetime)"`
	Updateat time.Time `orm:"type(datetime)"`
}

func (m *Hobby) TableName() string {
	return "t_hobby"
}

func (m *Hobby) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err != nil {
		return err
	}
	return nil
}


func (m *Hobby) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields ...);err != nil {
		return err
	}
	return nil
}

func (m *Hobby) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields ...);err != nil {
		return err
	}
	return nil
}

func (m *Hobby) Delete() error {
	if _,err := orm.NewOrm().Delete(m);err != nil {
		return err
	}
	return nil
}

func (m *Hobby) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}


func (m *Hobby) FindAll(container interface{}) bool{
	t,err := m.Query().All(container)
	if err == nil && t > 0 {
		return true
	}
	return false
}

func (m *Hobby) ToJsonString() string {
	data,_ := json.Marshal(m)
	return string(data)
}