package main

import "fmt"

var p = fmt.Println

var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1,intChan2}

var numbers = []int {1,2,3,4,5}

func main(){
	select {
	case getChan(0) <-  getNumber(0):
		p("The 1th case is selected")
	case getChan(1) <- getNumber(1):
		p("The 2th case is selected")
	default:
		p("Default")
	}
}

func getNumber(i int) int {
	p("numbers[%d] \n",i)
	return numbers[i]
}

func getChan(i int) chan int {
	p("channels [%d] \n",i)
	return channels[i]
}
