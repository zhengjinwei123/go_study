package main

import (
	"os"
	"fmt"
)

func main(){

	f := createFile("G://zjw_go.txt")
	defer closeFile(f)
	writeFile(f)

	var content []byte
	n,err := f.Read(content)
	if err != nil{
		panic(err)
	}
	fmt.Println("read n:",n,content)
}

func createFile(p string) *os.File{
	fmt.Println("creating")
	f,err := os.Create(p)
	if err != nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("writing")
	fmt.Fprintln(f,"data")
}

func closeFile(f *os.File){
	fmt.Println("closing")
	f.Close()
}