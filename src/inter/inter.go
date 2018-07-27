package inter

import "fmt"

type Animal interface {
	Eat()
	Talk()
}

type Dog struct {
	Name string
}

func (d *Dog) Eat(){
	fmt.Println("dog is eating")
}

func (d *Dog) Talk(){
	fmt.Println("dog is wawa")
}

type Cat struct {

}

func (d *Cat) Eat(){
	fmt.Println("cat is wawa")
}

func (d *Cat) Talk(){
	fmt.Println("cat is miaomiao")
}



func TestInterface(){
	var a Animal

	var d Dog
	d.Eat()

	a = &d
	a.Eat()

	var c Cat
	a = &c
	a.Eat()
}