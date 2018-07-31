package main

import "fmt"

func main(){
	dataChan := make(chan int,5)
	syncChan1 := make(chan struct{},1)
	syncChan2 := make(chan struct{},2)

	go func() {
		<- syncChan1
		for{
			if elem,ok := <- dataChan;ok{
				fmt.Println("Received:%d [received] \n",elem)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for i := 0; i < 5; i++{
			dataChan <- i
			fmt.Println("Sent:%d [sender] \n",i)
		}
		close(dataChan)

		syncChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()

	<- syncChan2
	<- syncChan2
}