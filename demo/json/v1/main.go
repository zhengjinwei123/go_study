package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIp string
}

type Serverslice struct {
	Servers []Server
}

func main(){
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomes","Mortical"]}`)
	var f interface{}
	err := json.Unmarshal(b,&f)

	if err != nil {
		panic(err)
	}
	 m := f.(map[string]interface{})

	 for k,v := range m{
	 	switch vv := v.(type){
		case string:
			fmt.Println(k,"is string ",vv)
		case int:
			fmt.Println(k, "is int ",vv)
		case float64:
			fmt.Println(k,"is float64",vv)
		case []interface{}:
			fmt.Println(k,"is an array:")
		    for i,u := range vv{
		    	fmt.Println(i,u)
			}
		default:
			fmt.Println(k," is of a type I don't know how to handle")
		}
	 }
}
