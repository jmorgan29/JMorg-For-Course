package main


//The prime factors of 13195 are 5,7,13,and 29. What is the largest prime factor of the number 600851475143
//cannot solve computationally with big number; start with 13195
import ("fmt"
"math"
)

func main() {
		//y,n:=primelist(50000)
		//fmt.Println(y)
		//fmt.Println(n)

	//x1 := []int{1,1,3,9} //error occurs when the last one is taken out
	//This function is a mess; use a boolean operator,initialized as true, looping over the values of x1. If a value from x2 is equal to the main one
	//set the boolean equal to false. If at the end of the loop the boolean remains true, add the value to another (initially empty) slice
	//x2 := []int{1,7,3}
	//y2:=common_removal(x1,x2)
	//fmt.Println(y2)
	xin:=20142987923
	y3:=factorization(xin)
	//can be used to factor 1000000000 (10^9)
    largestfactor:=y3[len(y3)-1]
    if largestfactor==0 {
    	fmt.Println("Failure to solve, pending future code improvement")
    } else {
    	fmt.Println(y3)
    	fmt.Println("The largest prime factor of", xin, "is", largestfactor)}

    x10,x11:=primelist(3759)
    fmt.Println(x10)
    fmt.Println(x11)

}

func  factorization(x int) ([]int) {
	factors:=make([]int,0,2*x)
	xuse:=x
	for i:=0;i<2*x;i++ {
		y1,y2:=addnumbers(xuse)
		factors=append(factors,y1)
		xuse=y2
		if y2==1 {
			break
		}
	}

	return factors

}

func addnumbers(x int) (y1 int, y2 int) {
	xtake:=float64(x)
	for i:=0;i<50;i++{
		if xtake<50000 {break}
		xtake=math.Sqrt(xtake)
	}
	xtake2:=int(xtake)
	y,_:=primelist(xtake2)
	for i:=0;i<len(y);i++ {
		if x%y[i]==0 {
			y1 = y[i]
			y2 = x / y[i]
			break
		} else {   //else signifies failure to solve
			y1=0
			y2=0
		}

	}
	return y1,y2
}

	func primelist(x int) ([]int,int) {
		numb := make([]int, 0, 2*x)
		//maxval:=float64(x)
		count:=0
		for i := 2; i < x+1; i++ {
			numb = append(numb, i)			}

		for j:= 2; j <len(numb); j++ {
			act := numb[count]
			fact := make([]int, 0, 2*x)
			for k := 2; k*act <= x; k++ {
				fact = append(fact, k*act)
				//numb = common_removal(numb, fact)
			}
			numb=common_removal(numb,fact)
			count=count+1
		}

	return numb,len(numb)
	}

	//remove from x1 if it is included in x2
	func common_removal(x1 []int, x2 []int) (y [] int){
		var status bool
	  y=make([]int,0,len(x1))
		for i:=0;i<len(x1);i++ {
			status=true
		for j:=0;j<len(x2);j++{
				if x2[j]==x1[i] {
					status = false
					break
				}
			}
			if status==true {
				y=append(y,x1[i])
			}
		}
		return y
	}



