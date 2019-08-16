package main

import "fmt"

var x string
func main() {
	fmt.Printf("What's your name?>")
	fmt.Scan(&x)
	fmt.Println("Hello",x)
}
