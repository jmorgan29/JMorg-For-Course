package main

import "fmt"

func main() {
	a:=foo(1,2)
	fmt.Println(a)
	b:=foo(1,2,3)
	fmt.Println(b)
	aSlice:=[]int{1,2,3,4}
	c:=foo(aSlice...)
	fmt.Println(c)
	d:=foo()
	fmt.Println(d)
}

func foo(m...int) int {
	//fmt.Println(m)
	//fmt.Printf("%T \n",m)
	var curr int
	for _, v:=range m{
		if v>curr {
			curr=v }
	}
	return curr
}