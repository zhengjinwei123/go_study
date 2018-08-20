package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	Username string
	Sex string
	Age int
	AvatarUrl string
}

func NewUser(username string,sex string,age int,avatar string) *User{
	user := new(User)
	user.Username = username
	user.Sex = sex
	user.Age = age
	user.AvatarUrl = avatar
	return user
}

func test1(){
	u := NewUser("user1","女",18,"xxx.jpg")
	fmt.Printf("user = %#v \n",u)
}

type User2 struct {
	Username string `json:"xxxx"`
	Sex string `json:"sex"`
	Score float32 `json:"score"`
	Age int32 `json:"age"`
}

func test2(){
	u := &User2{
		Username: "user01",
		Sex: "男",
		Score: 99.2,
		Age : 18,
	}

	data,_ := json.Marshal(u)
	fmt.Printf("json str:%s \n",string(data))
}

func test3(){
	type User struct {
		Username string
		Sex string
		Age int
		AvatarUrl string
	}

	var user *User
	fmt.Printf("user = %v \n",user)
	var user01 *User = &User{}

	user01.Age = 18
	user01.Username = "user01"
	fmt.Printf("user01 = %#v \n",user01)

	var user02 *User = &User{
		Username: "user02",
		Age: 18,
	}

	fmt.Printf("user02=%#v \n",user02)
	var user03 *User = new(User)
	user03.Username = "user03"
	user03.Age = 18

	fmt.Printf("user03= %#v \n",user03)
}

func test4(){
	type Address struct{
		Province string
		City string
	}

	type User struct {
		Username string
		Sex string
		address *Address
	}

	user := &User{
		Username:"user01",
		Sex:"man",
		address:&Address{
			Province:"beijing",
			City:"beijing",
		},
	}
	fmt.Printf("user=%#v \n",user)
}

var (
	coins = 100
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int,len(users))
)
func calcCoin(username string) int{
	var sum int = 0
	for _,char := range username {
		switch char {
		case 'a','A':
			sum += 1
		case 'e','E':
			sum += 1
		case 'i','I':
			sum += 2
		case 'o','O':
			sum += 3
		case 'u','U':
			sum += 5
		}
	}

	return sum
}

func dispatchCoin() int{
	var left int = coins
	for _,username := range users{
		allCoin := calcCoin(username)
		left = left - allCoin
		value,ok := distribution[username]
		if !ok{
			distribution[username] = allCoin
		} else {
			distribution[username] = value + allCoin
		}
	}
	return left
}

func test5(){
	left := dispatchCoin()
	for username,coin := range  distribution{
		fmt.Printf("user:%s have %d coins \n",username,coin)
	}
	fmt.Printf("left coin:%d \n",left)
}

func main(){
	test1()
	test2()
	test3()
	test4()
	test5()
}
