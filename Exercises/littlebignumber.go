package main

import "fmt"

var a int
var b int
func main() {
	fmt.Printf("Enter an integer>")
	fmt.Scan(&a)
	fmt.Printf("Enter another integer>")
	fmt.Scan(&b)
	if a>=b {
		fmt.Println(a%b)
	} else {
		fmt.Println(b%a)
	}
}
