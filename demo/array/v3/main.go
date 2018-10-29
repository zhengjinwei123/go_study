package  main

import (
	"strconv"
	"fmt"
)

type Element interface {

}

type List []Element

type Person struct {
	name string
	age int
}

func (p Person) String() string{
	return "(name:"+p.name+" - age :"+strconv.Itoa(p.age)+" years)"
}

func main(){
	list := make(List,3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{
		"Denis",
		50,
	}

	for index,element := range list {
		switch value := element.(type){
		case int:
			fmt.Printf("list[%d] is an int and its value is %d \n",index,value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s \n",index,value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s \n",index,value)
		default:
			fmt.Printf("list[%d] is a different type \n",index)
		}
	}
}

