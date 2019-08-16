package main

import "fmt"

func main() {
	n:=maximum(94,20,30,60,70,34,41,93.6,-7,94)
	fmt.Println(n)
}

func maximum(m...float64) float64 {
	fmt.Println(m)
	fmt.Printf("%T \n",m)
	var curr float64
	for _, v:=range m{
		if v>curr {
			curr=v }
	}
	return curr
}