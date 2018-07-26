package main

import (
	"mypath"
	"fmt"
	"goroute"
	"time"
	"math/rand"
	"unicode/utf8"
	"sort"
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

	var s3 [] int
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
}

func testMapSlice(){
	s := make([]map[string]int,10)

	for i := 0;i < len(s); i++{
		s[i] = make(map[string]int,100)
	}

	s[0]["abc"] = 100
	s[5]["abc"] = 200

	fmt.Println(s)
}

func testMap(){
	var a map[string]int
	a = make(map[string]int,1000)
	a["a"] = 0
	a["b"] = 12
	a["cd"] = 45

	fmt.Println(a)

	var keys []string
	for k,v := range a {
		fmt.Printf("a[%s] = %d \n",k,v)
		keys = append(keys,k)
	}

	fmt.Println("")

	sort.Strings(keys)
	for _,k := range keys{
		fmt.Printf("a[%s] = %d \n",k,a[k])
	}


	val,exists := a["cd"]
	fmt.Printf("val=%d exist= %t \n",val,exists)

	if exists{
		fmt.Printf("val = %d \n",val)
	} else{
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
	fmt.Printf("%p \n",&b)
	fmt.Printf("%p \n",b)
	fmt.Printf("%p \n",a)
	fmt.Printf("%p \n",&a)

	b = &a
	fmt.Printf("%d \n",*b)

	*b = 100
	fmt.Printf("a = %d \n",a)

	fmt.Println("---------test1 end-------------------")

	testClosure()
}

func Adder() func(int) int{
	var x int
	f := func(i int) int {
		x = x +i
		return x
	}

	return f
}

func testClosure(){
	f1 := Adder()

	fmt.Println(f1(10))
	fmt.Println(f1(20))
	fmt.Println(f1(30))
}