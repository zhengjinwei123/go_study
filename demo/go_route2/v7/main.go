package main

import "fmt"

func main(){
	chanCap := 5
	intChan := make(chan int,chanCap)

	for i := 0;i < chanCap;i++{
		select {
		case intChan <- 1:
			fmt.Println("1")
	    case intChan <- 2:
			fmt.Println("2")
	    case intChan <- 3:
			fmt.Println("3")
		}
	}

	for i := 0;i < chanCap;i++{
		fmt.Printf("haha %d \n",<- intChan)
	}
}