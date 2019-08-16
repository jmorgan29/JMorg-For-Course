package main

import "fmt"

var xnew int
func main() {
	y:=fib(49)
	fmt.Println(y)
}

func fib(x int) (int) {
	if x<=2 {
		return 0
	}
	x1:=1
	x2:=2
	//Deal with numbers 2 and below!!
	num_sum:=2
	for {
		xnew=x1+x2
		x1=x2
		x2=xnew
		if xnew>=x {break}
		if xnew%2==0 {num_sum=num_sum+xnew}
	}

	return num_sum
}

//This program returns the sum of the even Fibronacci numbers that are lower than some integer (4000000 in this example)
//The Fibronacci numbers are defined by the infinite sequence (1,2,3,5,8,...) where each number is the sum of the two previous numbers in the sequence