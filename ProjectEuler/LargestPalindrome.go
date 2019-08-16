package main

import ("fmt"
"strconv"
"github.com/GoesToEleven/GolangTraining/02_package/stringutil"
"math"
)


//Try factoring first with the prime numbers drawn from a set to get a smaller number

//A palindromic number reads the same both ways. The largest palindrome made from the product of 2-digit numbers is
//91*99. Find the largest palindrome made from the product of two 3-digit numbers

func main() {
	minval:=100
	maxval:=999
	var y string
	var yrev string
	var ynew int
	var finalval int
	palindrome_list := make([]int, 10, 100)   //arbitrary size
 	for i:=minval; i<maxval+1; i++ {
 		//fmt.Println(i)
		for j:=minval; j<maxval+1; j++{
			y=strconv.Itoa(i*j)
			yrev=stringutil.Reverse(y)
			ynew,_=strconv.Atoi(yrev)
			if i*j==ynew {palindrome_list = append(palindrome_list,i*j)}
		}
	}

	//find maximum number in palindrome_list
	for _,n:=range palindrome_list {
		finalval=int(math.Max(float64(finalval),float64(n)))
	}

	//fmt.Println(finalval)
	fmt.Printf("The largest palindrome made from the product of two three digit numbers is %v\n.", finalval)
}
