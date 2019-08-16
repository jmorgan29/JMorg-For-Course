package main

import ("fmt"
"math"
)

const(
	a=3
	b=2
	c=1
	d=-2
	e=-7
	tol=1e-10
	maxit=50
)



func main() {
	fmt.Println("i","x","y","y'")
	x:=20.0
	for i:=1;i<maxit;i++{

		y:=f(x)
		yprime:=fprime(x)
		z:=x-y/yprime
		fmt.Printf("%v\t",i)
	fmt.Printf("%v\t", x)

	fmt.Printf("%v\t", y)
	fmt.Printf("%v\n",yprime)

	if y<tol {break}
	x=z
	}

	//var add*float64=&x
	//fmt.Println(add)
	//fmt.Printf("%v\n",add)
	//*add=x-y/yprime
	//fmt.Printf(x,y,yprime)
}

func f(x float64) float64{
	y:=a*math.Pow(x,a)+d*math.Pow(x,b)+c*x+e
	return y
	//3*x^3-2*x^2+x-7
}

func fprime(x float64) float64{
	yprime:=a*a*math.Pow(x,a-1)+d*b*math.Pow(x,b-1)+c
	return yprime
}

//fmt.Printf("%b\t", TB) indicates to leave a space
//fmt.Printf("%d\n", TB) indicates to start a new line