package main

import "fmt"

func testString(){
	var str = "hello"

	fmt.Printf("str[0]=%c len(str)=%d \n",str[0],len(str))

	for index,val := range str{
		fmt.Printf("str[%d]=%c\n",index,val)
	}

	var byteSlice []byte
	byteSlice = []byte(str)
	byteSlice[0] = '0'
	str = string(byteSlice)

	fmt.Printf("after modify:%s \n",str)

	fmt.Printf("len(str) = %d \n",len(str))

	str = "hello,zjw"
	fmt.Printf("len(str) = %d \n",len(str))

	str = "中间123"
	fmt.Printf("last:len(str)=%d \n",len(str))

	var b rune = '中'
	fmt.Printf("b=%c\n",b)

	var runeSlice []rune
	runeSlice = []rune(str)
	fmt.Printf("str 长度:%d,len(str)=%d\n",len(runeSlice),len(str))
}

func testReverseStringV1(){
	var str = "hello中文"
	var bytes []rune = []rune(str)

	for i := 0; i < len(bytes) / 2;i++{
		tmp := bytes[len(bytes) - i - 1]
		bytes[len(bytes)-i-1] = bytes[i]
		bytes[i] = tmp
	}

	str = string(bytes)
	fmt.Println(str)
}

func testHuiWen(){
	var str = "上海自来水来自海上"
	var r []rune = []rune(str)

	for i := 0; i < len(r) /2;i++{
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}

	str2 := string(r)
	if str2 == str{
		fmt.Println(str," is huiwen")
	}else{
		fmt.Println(str," is not huiwen")
	}
}


func main(){
	testString()
	testReverseStringV1()
	testHuiWen()
}