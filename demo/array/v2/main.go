package  main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func testArray1(){
	var a [5]int
	a[0] = 200
	a[1] = 300
	fmt.Print(a)
}

func testArray2(){
	var a [5]int =  [5]int{1,2,3,4,5}
	fmt.Println(a)
}

func testArray3(){
	a := [5]int{1,2,3,4,5}
	fmt.Println(a)
}

func testPoint1(){
	var a int32
	a = 1000
	fmt.Printf("the address of a:%p,a:%d \n",&a,a)

	var b *int32
	fmt.Printf("the address of b:%p,b:%v \n",&b,b)

	if b == nil{
		fmt.Println("b is nil address")
	}

	b = &a
	fmt.Printf("the addr of b:%p,b:%v \n",&b,b)
}

func testMap1(){
	var a map[string] int
	fmt.Printf("a:%v \n",a)

	if a == nil {
		a = make(map[string]int,16)
		fmt.Printf("a=%v \n",a)
		a["stu1"] = 1000
		a["stu2"] = 2000
		fmt.Printf("a = %#v \n",a)
	}
}

func testMap2(){
	rand.Seed(time.Now().Unix())
	var a map[string]int = make(map[string]int,1024)
	for i:= 0;i < 128;i++{
		key := fmt.Sprintf("stu%d",i)
		value := rand.Intn(1000)
		a[key] = value
	}

	var keys []string = make([]string,0,128)
	for key,value := range a{
		fmt.Printf("map[%s] = %d \n",key,value)
		keys = append(keys,key)
	}

	sort.Strings(keys)

	for _,value := range keys {
		fmt.Printf("key:%s val:%d \n",value,a[value])
	}
}

func testMap3(){
	rand.Seed(time.Now().Unix())
	var s []map[string]int
	s = make([]map[string]int,5,16)
	for index,val := range s {
		fmt.Printf("slice[%d] = %v\n",index,val)
	}

	fmt.Println()

	s[0] = make(map[string]int,16)
	s[0]["stu01"] = 1
	s[0]["stu02"] = 2
	s[0]["stu03"] = 3

	for index,val := range s{
		fmt.Printf("slice[%d] = %v \n",index,val)
	}
}

func testMap4(){
	rand.Seed(time.Now().Unix())
	var s map[string][]int

	s = make(map[string][]int,16)
	key := "stu01"

	value,ok := s[key]

	if !ok{
		s[key] = make([]int,0,16)
		value = s[key]
	}

	value = append(value,100)
	value = append(value,200)
	value = append(value, 300)
	s[key] = value

	fmt.Printf("map:%v \n",s)
}

func testMap5(){
	var a map[string]int = map[string]int {
		"stu01" : 1000,
		"stu02" : 2000,
		"stu03" : 3000,
	}

	fmt.Printf("%#v \n",a)

	a["stu01"] = 88888
	a["stu04"] = 8989

	var key string = "stu04"
	fmt.Printf("the value of key[%s] is %d \n",key,a[key])
}

func testMap6(){
	var a map[string]int
	a = make(map[string]int,16)

	a["stu01"] = 1
	a["stu02"] = 2
	a["stu03"] = 3

	fmt.Printf("a = %#v \n",a)

	var result int
	var ok bool

	var key string = "stu06"

	result,ok = a[key]
	if ok == false {
		fmt.Printf("key %s is not exists \n",key)
	} else {
		fmt.Printf("key %s %d \n",key,result)
	}
}

func testMap7(){
	var a map[string]int

	fmt.Printf("----------------------\n a:%v \n",a)
	if a == nil {
		a = make(map[string]int,16)
		fmt.Printf("a = %v \n",a)

		a["stu01"] = 1000
		a["stu02"] = 1000
		a["stu03"] = 1000

		fmt.Printf("a = %#v \n",a)
		delete(a,"stu02")
		fmt.Printf("a = %#v \n",a)

		for key,_ := range a{
			delete(a,key)
		}

		fmt.Printf("after delete a = %#v \n",a)
	}
}

func main(){
	//testPoint1()
	//testMap1()
	//testMap2()
	//testMap3()

	testMap7()
}