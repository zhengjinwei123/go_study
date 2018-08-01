package main

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)


type Version struct {
	Id  int  `db:"id"`
	Released int `db:"released"`
	Major int `db:"major"`
	Minor int `db:"minor"`
	Text  string `db:"text"`
	ReleasedAt string `db:"releasedAt"`
}


var Db *sqlx.DB

func init(){
	database,err := sqlx.Open("mysql","root:123456@tcp(10.18.97.143:3306)/zjw_rob")
	if err != nil{
		Db = nil
		fmt.Println("open mysql failed,",err)
	}

	Db = database
}

func query(){
	var versions []Version
	err := Db.Select(&versions,"select id,released,major,minor,text,releasedAt from Version")
	if err != nil{
		fmt.Println("select error:",err)
		return
	}

	fmt.Printf("%#v \n",versions)
}

func insert(){
	r,err := Db.Exec("insert into Version(id,released,major,minor,text,releasedAt) values(?,?,?,?,?,?)",100,1,1,100,"zjw","2018-08-01 18:09-01")

	if err != nil{
		fmt.Println("exec insert err:",err)
		return
	}

	id ,err := r.LastInsertId()
	if err != nil{
		fmt.Println("exec failed,",err)
		return
	}
	fmt.Println("insert success,id:",id)
}

func update(){
	_,err := Db.Exec("update Version set text=? where id=?","zjw1",100)
	if err != nil{
		fmt.Println("exec update failed,",err)
		return
	}
	fmt.Println("exec update success")
}

func delete(){
	_,err := Db.Exec("delete from Version where id=?",100)
	if err != nil{
		fmt.Println("exec delete failed,",err)
		return
	}
	fmt.Println("exec delete success")
}

func main(){
	fmt.Println(Db)

	if Db == nil{
		fmt.Println("db is nil")
		os.Exit(1)
	}

	query()

	insert()

	update()

	delete()
}