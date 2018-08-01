package main

import (
	"os"
	"fmt"
)

func main(){
	//panic("a problem")

	_,err := os.Create("/tmp/file")
	if err != nil{
		panic(fmt.Sprintf("err:%s",err))
	}
}