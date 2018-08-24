package main

import (
	"net/http"
	"fmt"
	"strings"
	"html/template"
	"log"
	"path/filepath"
)

func sayHelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("schema",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k,v := range r.Form{
		fmt.Println("key:",k)
		fmt.Println("value:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello zjw")
}

func login(w http.ResponseWriter,r *http.Request){
	fmt.Println("method:", r.Method) //获取请求的方法
	currentDir,err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentDir)

	tplPath := currentDir + "/demo/array/v6/login.gtpl"
	if r.Method == "GET" {
		t, _ := template.ParseFiles(tplPath)
		t.Execute(w, nil)
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Fprintf(w,"login success")
	}
}


func main(){
	http.HandleFunc("/",sayHelloName)
	http.HandleFunc("/login",login)

	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}
