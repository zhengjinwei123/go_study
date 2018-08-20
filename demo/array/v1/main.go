package main

import (
	"math/rand"
	"time"
	"fmt"
)

func sumArray(a [10]int) int{
	var sum int =0
	for i := 0; i < len(a);i++{
		sum = sum + a[i]
	}
	return sum
}

func testArraySum(){
	rand.Seed(time.Now().Unix())
	var b [10]int

	for i := 0; i< len(b);i++{
		b[i] = rand.Intn(1000)
	}
	sum := sumArray(b)
	fmt.Printf("sum=%d\n",sum)
}

func TwoSum(a [5]int,target int){
	for i := 0;i < len(a);i++{
		other := target - a[i]
		for j := i+1 ;j < len(a);j++{
			if a[j] == other{
				fmt.Printf("(%d,%d)\n",i,j)
			}
		}
	}
}


func testTwoSum() {

	//var b [5]int = [5]int{1,3,5,8,7}
	//b := [5]int{1,3,5,8,7}
	b := [...]int{1, 3, 5, 8, 7}
	TwoSum(b, 8)
}

func main() {
	//testArraySum()
	testTwoSum()
}