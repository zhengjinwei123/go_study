package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/orm"
	"mygoblog/util"
	"strings"
	"time"
	"goblog/models"
)

type Post struct {
	Id         int
	UserId     int       `orm:"index"`
	Author     string    `orm:"size(15)"`
	Title      string    `orm:"size(100)"`
	Color      string    `orm:"size(7)"`
	UrlName    string    `orm:"size(100);index"`
	UrlType    int8
	Content    string    `orm:"type(text)"`
	Tags       string    `orm:"size(100)"`
	PostTime   time.Time `orm:"type(datetime);index"`
	Views      int
	Status     int8
	UpdateTime time.Time `orm:"type(datetime)"`
	IsTop      int8
}

func (m *Post) TableName() string {
	return TableName("post")
}

func (m *Post) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Post) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields ...); err != nil {
		return err
	}
	return nil
}

func (m *Post) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields ...); err != nil {
		return err
	}
	return nil
}

func (m *Post) Delete() error {
	if m.Tags != "" {
		o := orm.NewOrm()
		oldTags := strings.Split(strings.Trim(m.Tags,","),",")
		o.QueryTable(&Tag{}).Filter("name_in",oldTags).Update(orm.Params{"count":orm.ColValue(orm.ColMinus,1)})
		o.QueryTable(&TagPost{}).Filter("postid",m.Id).Delete()
	}

	if _,err := orm.NewOrm().Delete(m);err != nil {
		return err
	}
	return nil
}

func (m *Post) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Post) ColorTitle() string {
	if m.Color != "" {
		return fmt.Sprintf("<span style=\"color:%s\">%s</span>",m.Color,m.Title)
	} else {
		return m.Title
	}
}

func (m *Post) Link() string {
	if m.UrlName != "" {
		if m.UrlType == 1 {
			return fmt.Sprintf("/%s",util.Rawurlencode(m.UrlName))
		}
		return fmt.Sprintf("/article/%s",util.Rawurlencode(m.UrlName))
	}
	return fmt.Sprintf("/article/%d",m.Id)
}

func (m *Post) TagsLink() string {
	if m.Tags == "" {
		return ""
	}

	var buf bytes.Buffer

	arr := strings.Split(strings.Trim(m.Tags,","),",")
	for k,v := range arr {
		if k > 0 {
			buf.WriteString(", ")
		}
		tag := Tag{Name:v}
		buf.WriteString(tag.Link())
	}
	return buf.String()
}

func (m *Post) Except() string {
	if i := strings.Index(m.Content,"_ueditor_page_break_tag_"); i != -1 {
		return m.Content[:i]
	}
	return m.Content
}