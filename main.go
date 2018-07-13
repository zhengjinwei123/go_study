package main

import (
	"fmt"
	"mypath"
	"errors"
)

func divide(divident int, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("devision by zero")
		return
	}
	result = divident / divisor
	return
}

func main() {
	fmt.Println("hello world")
	x := mypath.Sqrt(100)
	fmt.Println(x)

	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[4])

	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 6}
	fmt.Println(b)

	var two [2][3]int;
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			two[i][j] = i + j
		}
	}

	fmt.Println(two)

	// 切片
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)

	fmt.Println(len(s))

	s = append(s, "d")
	s = append(s, "e", "f", "g")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println(l)

	l1 := s[:5]
	fmt.Println(l1)

	l2 := []string{"g", "h", "j"}
	fmt.Println(l2)

	l3 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		l3[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			l3[i][j] = j + j
		}
	}
	fmt.Println(l3)

	// 关联数组
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println(m["k1"])

	delete(m, "k1")
	fmt.Println(m)

	jj, prs := m["k2"]
	fmt.Println("prs", prs, jj)

	n := map[string]int{"f": 1, "g": 2}
	fmt.Println(n)

	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for i, n := range nums {
		fmt.Println(i, n)
		sum += n
	}

	fmt.Println(sum)

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs {
		fmt.Printf("%s->%s \r\n",k,v)
	}

	for i,c := range "go" {
		fmt.Println(i,c)
	}
}
