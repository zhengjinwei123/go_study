package main

import (
	"net/http"
	"fmt"
	"strings"
	"html/template"
	"log"
	"path/filepath"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello zjw")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	currentDir, err := filepath.Abs("./")
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
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Fprintf(w, "login success")
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	currentDir, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}

	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))


		t, _ := template.ParseFiles(currentDir + "/demo/array/v6/upload.gtpl")
		t.Execute(w, token)
	} else {
		fmt.Println("hahahhaha")
		maxMemory := 32 << 20
		fmt.Println("maxMemory:",maxMemory)
		r.ParseMultipartForm(int64(maxMemory))
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintf(w,"%v",handler.Header)

		f,err := os.OpenFile("./"+
			handler.Filename,os.O_RDWR | os.O_CREATE,0666)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f,file)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
