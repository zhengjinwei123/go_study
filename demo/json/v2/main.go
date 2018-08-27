package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP string `json:"serverIp"`
	ID int `json:"id,omitempty,"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main(){
	var s Serverslice

	s.Servers = append(s.Servers,Server{ServerName:"Shanghai_VPN",ServerIP:"127.0.0.1",ID:1})
	s.Servers = append(s.Servers,Server{ServerName:"Beijing_VPN",ServerIP:"127.0.0.2",ID:2})
	s.Servers = append(s.Servers,Server{ServerName:"Beijing_VPN",ServerIP:"127.0.0.2"})

	b,err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:",err)
	}
	fmt.Println(string(b))
}
