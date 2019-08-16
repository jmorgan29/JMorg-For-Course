package main

import "fmt"

//import "fmt"

func main() {
	fmt.Println(half(5))
}

func half(x int) (float64,bool){
	xp:=float64(x)
	//xp:=x
	y:=xp/2
	return y,x%2==0
}


