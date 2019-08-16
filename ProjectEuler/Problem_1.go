//This program was written to find the sum of the natural numbers below 1000 that are multiples of 3 or 5

//Will re-write in function form


package main

import "fmt"

var status bool
var multiples []int

func slice_sum(x []int) int{
	sum :=0
	for i :=range x {
		sum = sum+x[i]
	}
return sum
}


func main() {
a := []int{3,5}


for i:=1;i<1000;i++ {
	status=false
	for j := range a {
		if i%a[j] ==0 {
			status=true
			break}
	}
	if status==true {multiples=append(multiples,i)}
	}

fmt.Println(multiples)
fmt.Println(slice_sum(multiples))
}


