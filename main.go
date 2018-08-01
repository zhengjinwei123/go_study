package main

import (
	"encoding/json"
	"fmt"
	"goroute"
	"inter"
	"math/rand"
	"mypath"
	"sort"
	"time"
	"unicode/utf8"
)

func modify(a int) {
	a = 10
	return
}

func modify1(a *int) {
	*a = 10
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	return
}

func swap1(a int, b int) (int, int) {
	return b, a
}

const (
	Man    = 1
	Female = 2
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Sum(b []int) int {
	var sum int
	cnt := len(b)
	for i := 0; i < cnt; i++ {
		sum = sum + b[i]
	}

	return sum
}

func testSliceCap() {
	a := make([]int, 5, 10)
	a[0] = 11
	a[1] = 12
	a[2] = 13

	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a[4] = 100
	b := a[1:3]

	fmt.Println("a=%#v,len(a) = %d,cap(a) = %d \n", a, len(a), cap(a))
	fmt.Println("b=%#v,len(b) = %d,cap(b) = %d \n", b, len(b), cap(b))
}

func main() {
	inter.TestInterface()

	sum, _ := mypath.Add(100, 300)
	fmt.Println(sum)

	var pipe chan int
	pipe = make(chan int, 1)
	go goroute.Add(1, 100, pipe)

	sum2 := <-pipe

	fmt.Println(sum2)

	a := 5
	b := make(chan int, 1)

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	modify(a)
	fmt.Println("a=", a)
	modify1(&a)
	fmt.Println("a=", a)

	first := 100
	second := 200

	swap(&first, &second)

	fmt.Println(first, second)

	first, second = swap1(first, second)

	fmt.Println(first, second)

	first, second = second, first

	fmt.Println(first, second)

	fmt.Println("------------------------------")

	now := time.Now()
	second1 := now.Unix()

	if second1%Female == 0 {
		fmt.Println("female")
	} else {
		fmt.Println("man")
	}

	fmt.Println("---------rand.Int")
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	fmt.Println("---------rand.Intn")
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}

	fmt.Println("---------rand.float32")
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}

	//var arr [5] int = [5]int {1,2,3,4,5}

	var arr1 = [5]int{1, 2, 3}

	for k, v := range arr1 {
		fmt.Println(k, v)
	}

	var arr2 = [5]int{1}
	fmt.Println(len(arr2))

	var arr3 = [5]string{1: `zjw`, 3: "jahaaa"}

	for i := 0; i < len(arr3); i++ {
		fmt.Println(i, arr3[i])
	}

	for k, v := range arr3 {
		fmt.Println(k, v)
	}

	var arr4 = [...]int{1, 3, 2, 3}
	arr4[0] = 2
	arr4[1] = 3
	arr4[2] = 4
	for k, v := range arr4 {
		fmt.Println(k, v)
	}

	sl := make([]int, 10)
	sl[0] = 1
	sl[1] = 2
	sl[2] = 3
	sl[3] = 4

	fmt.Println("sum slice", Sum(sl), len(sl))

	testSliceCap()

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 10)

	copy(s2, s1)

	fmt.Println(s2)

	var s3 []int
	s3 = make([]int, 5)
	var s4 []int = []int{10, 11, 13, 24}

	s3 = append(s3, 1, 1, 2, 3, 4, 5, 6)
	s3 = append(s3, s4...)
	fmt.Printf("zjw s3:%#v \n", s3)

	//-----------字符串---------------
	var str = "hello world"
	var bs []byte = []byte(str)
	bs[0] = 'a'
	str1 := string(bs)

	fmt.Println(str1)

	str11 := "hello world"

	b11 := []byte(str11)

	for i := 0; i < len(b11)/2; i++ {
		b11[i], b11[len(b11)-i-1] = b11[len(b11)-i-1], b11[i]
	}

	fmt.Println(string(b11))

	str12 := "hello wolrd我们爱中国"
	b12 := []rune(str12)

	for i := 0; i < len(b12)/2; i++ {
		b12[i], b12[len(b12)-i-1] = b12[len(b12)-i-1], b12[i]
	}

	str13 := string(b12)
	fmt.Println(str13)

	fmt.Printf("len(str12) = %d,len(rune) = %d ,len(char count) = %d \n", len(str12), len(str13), len([]rune(str12)))

	fmt.Printf("runeCountInString = %d", utf8.RuneCountInString(str12))

	for {

		start := time.Now().UnixNano()
		now := time.Now()
		fmt.Printf("type of now is:%T \n", now)

		year := now.Year()
		month := now.Month()
		day := now.Day()

		str := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d \n", year, month, day, now.Hour(), now.Minute(), now.Second())

		fmt.Println(str)

		fmt.Printf("timestamp:%d \n", now.Unix())

		fmt.Println("zjw", now.Format("2006-01-02 03:04:05"))

		time.Sleep(time.Second)

		end := time.Now().UnixNano()

		cost := (((end - start) / 1000) / 1000) / 1000
		fmt.Printf("cost:%d s \n", cost)

		break
	}

	test1()

	testMap()

	testMapSlice()

	var i1 int
	var f1 float32
	var i32 int32
	justify(i1, f1, i32)
}

func justify(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case int:
			fmt.Printf("第 %d 个参数 is int\n", index)
		case int32:
			fmt.Printf("第 %d 个参数 is int32\n", index)
		case float32:
			fmt.Printf("第 %d 个参数 is float32\n", index)
		}
	}
}

func testMapSlice() {
	s := make([]map[string]int, 10)

	for i := 0; i < len(s); i++ {
		s[i] = make(map[string]int, 100)
	}

	s[0]["abc"] = 100
	s[5]["abc"] = 200

	fmt.Println(s)
}

func testMap() {
	var a map[string]int
	a = make(map[string]int, 1000)
	a["a"] = 0
	a["b"] = 12
	a["cd"] = 45

	fmt.Println(a)

	var keys []string
	for k, v := range a {
		fmt.Printf("a[%s] = %d \n", k, v)
		keys = append(keys, k)
	}

	fmt.Println("")

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("a[%s] = %d \n", k, a[k])
	}

	val, exists := a["cd"]
	fmt.Printf("val=%d exist= %t \n", val, exists)

	if exists {
		fmt.Printf("val = %d \n", val)
	} else {
		fmt.Println("now found")
	}

	val = a["aaaa"]

	fmt.Println(val)
}

func test1() {
	fmt.Println("---------test1 start-------------------")
	var a int
	a = 10

	fmt.Println(a)

	var b *int
	fmt.Printf("%p \n", &b)
	fmt.Printf("%p \n", b)
	fmt.Printf("%p \n", a)
	fmt.Printf("%p \n", &a)

	b = &a
	fmt.Printf("%d \n", *b)

	*b = 100
	fmt.Printf("a = %d \n", a)

	fmt.Println("---------test1 end-------------------")

	testClosure()

	testJson()

	testPeople()

	testMethod()
}

type People struct {
	Name string
	Age  int
}

func (s *People) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

func testMethod() {
	fmt.Println(" \n\n-------------------\n")
	var s People
	s.Set("zjw", 120)
	fmt.Println(s)
}

type Student1 struct {
	Score int
	People
	Name string
	int
}

func testPeople() {
	var s Student1
	s.Name = "zhw1"
	s.People.Name = "p111"
	s.Age = 200
	s.Score = 100
	s.int = 100

	fmt.Printf("hhhhahahah   %#v \n", s, s.int)
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:sex`
}

func testJson() {
	var a Student
	a.Age = 27
	a.Name = "zjw"
	a.Sex = "man"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("json marshal failed,err:%v \n", err)
		return
	}

	fmt.Printf("json data:%s \n", data)

	var s1 Student
	err = json.Unmarshal(data, &s1)
	if err != nil {
		fmt.Printf("json Unmarshal failed,err:%v \n", err)
		return
	}

	fmt.Printf("s1:%#v \n", s1)
}

func Adder() func(int) int {
	var x int
	f := func(i int) int {
		x = x + i
		return x
	}

	return f
}

func testClosure() {
	f1 := Adder()

	fmt.Println(f1(10))
	fmt.Println(f1(20))
	fmt.Println(f1(30))
}
