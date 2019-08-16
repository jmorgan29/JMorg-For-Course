package main

import("fmt"
	"math")
const(
	a=3
	b=2
	c=1
	d=-2
	e=-7
	tol=1e-10
	maxit=100
)

func main() {
	xp:=1.0
	xq:=5.0
	//fmt.Printf("%v\t", fprime(xq))
	if f(xp)*f(xq)>0{
		fmt.Println("Error: The zero is not contained between these points")
	} else {
		fmt.Println("i","xp","xq","yp","yq")
		for i:=1;i<maxit;i++{
			yp:=f(xp)
			yq:=f(xq)
			fmt.Printf("%v\t", i)
			fmt.Printf("%v\t",xp)
			fmt.Printf("%v\t",xq)
			fmt.Printf("%v\t",yp)
			fmt.Printf("%v\t",yq)
			xnew:=xq-f(xq)*(xq-xp)/(f(xq)-f(xp))
			fmt.Printf("%v\t",xnew)
			ynew:=f(xnew)
			fmt.Printf("%v\n",ynew)
			if math.Abs(ynew)<tol {break}
			if f(xp)*f(xnew)<0 {
				xq=xnew
			} else{
				xp=xnew
			}
		}
	}
}


func f(x float64) float64{
	y:=a*math.Pow(x,a)+d*math.Pow(x,b)+c*x+e
	return y
	//3*x^3-2*x^2+x-7
}

