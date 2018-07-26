package main

import (
	"mypath"
	"fmt"
	"goroute"
	"time"
	"math/rand"
)



func modify(a int){
	a = 10
	return
}

func modify1(a *int){
	*a = 10
}


func swap(a *int,b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	return
}

func swap1(a int, b int) (int,int){
	return b,a
}


const (
	Man = 1
	Female = 2
)

func init() {
	rand.Seed(time.Now().UnixNano())
}


func main() {
	sum,_ := mypath.Add(100,300)
	fmt.Println(sum)


	var pipe chan int
	pipe = make(chan int,1)
	go goroute.Add(1,100,pipe)

	sum2 := <- pipe

	fmt.Println(sum2)


	a := 5
	b := make(chan int,1)

	fmt.Println("a=",a)
	fmt.Println("b=",b)

	modify(a)
	fmt.Println("a=",a)
	modify1(&a)
	fmt.Println("a=",a)

	first := 100
	second := 200

	swap(&first,&second)

	fmt.Println(first,second)

	first,second = swap1(first,second)

	fmt.Println(first,second)

	first,second = second,first

	fmt.Println(first,second)

	fmt.Println("------------------------------")

	now := time.Now()
	second1 := now.Unix()

	if second1 % Female == 0{
		fmt.Println("female")
	}else{
		fmt.Println("man")
	}

	fmt.Println("---------rand.Int")
	for i:=0 ; i < 10; i++{
		a := rand.Int()
		fmt.Println(a)
	}

	fmt.Println("---------rand.Intn")
	for i :=0 ;i <10;i++{
		a := rand.Intn(100)
		fmt.Println(a)
	}

	fmt.Println("---------rand.float32")
	for i :=0 ;i <10;i++{
		a := rand.Float32()
		fmt.Println(a)
	}

	//var arr [5] int = [5]int {1,2,3,4,5}

	var arr1 = [5]int{1,2,3}

	for k,v := range arr1 {
		fmt.Println(k,v)
	}

	var arr2 = [5]int{1}
	fmt.Println(len(arr2))

	var arr3 = [5]string{1:`zjw`,3:"jahaaa"}


	for i := 0; i< len(arr3);i++{
		fmt.Println(i,arr3[i])
	}

	for k,v := range arr3 {
		fmt.Println(k,v)
	}

	var arr4 = [...]int{1,3,2,3}
	arr4[0] = 2
	arr4[1] = 3
	arr4[2] = 4
	for k,v := range arr4 {
		fmt.Println(k,v)
	}
}
