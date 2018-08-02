package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main(){
	c,err := redis.Dial("tcp","10.18.97.143:6379")
	if err != nil{
		fmt.Println("Connect to redis err:",err)
		return
	}
	defer c.Close()

	c.Do("AUTH","123456")
	c.Do("SET","pzpz",123)
	a,err := redis.Int(c.Do("GET","pzpz"))
	if err != nil{
		panic("set redis error")
	}

	fmt.Println(a)
}