package main

import (
	"net/http"
	"fmt"
)

type MyMux struct{

}

func (p *MyMux) ServeHTTP(w http.ResponseWriter,r *http.Request){
	if r.URL.Path == "/" {
		sayHelloName(w,r)
		return
	}

	http.NotFound(w,r)
	return
}

func sayHelloName(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello zjw")
}

func main(){
	mux := &MyMux{}
	http.ListenAndServe(":9999",mux)

}
