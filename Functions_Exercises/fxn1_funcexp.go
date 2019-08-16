package main

import "fmt"


//Need to expand to more variables
//import "fmt"
func half(x int) func(int) (int,bool) {
	return func(int) (int,bool) {
		return x/2,x%2==0
	}
}


func main() {




	x:=20
	fxn:=half(x)
	fmt.Println(fxn(x))
	//fmt.Printf("%T\n", fxn)
}