
//NBA Example created by JCM on 5/5/2019
package main

import (
	"fmt"
	"sort"
)

type BasketballTeam []string

func (b BasketballTeam) Len() int {return len(b)}
func (b BasketballTeam) Less(i,j int) bool {return b[i]<b[j]}
func (b BasketballTeam) Swap(i,j int) {b[i],b[j]=b[j],b[i]}

func (b BasketballTeam) Sort() {
	sort.Sort(b)
}

func (b BasketballTeam) SortReverse(){
	b.Sort()
	length:=b.Len()
	arg1:=0
	arg2:=0
	for i:=0; i<=length/2-1; i++ {
	//	fmt.Println(i)
		arg1,arg2=i,length-1-i
		b.Swap(arg1,arg2)
	}
}
func main() {
	Rockets:=BasketballTeam{"Harden","Capella","Faried","Gordon","Tucker","Nene","Paul","Rivers","Shumpert"}
	Knicks:=BasketballTeam{"Knox","Kornet","Robinson","Dotson","Hezonja","Mudiay"}
	fmt.Println(Rockets)
	Rockets.Swap(1,3)
	fmt.Println(Rockets)
	Rockets.Sort()
	fmt.Println(Rockets)
	Rockets.SortReverse()
	fmt.Println(Rockets)
	// Swtich To Knicks Example
	Knicks.Sort()
	fmt.Println(Knicks)
	Knicks.SortReverse()
	fmt.Println(Knicks)

}




