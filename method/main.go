package main 

import "fmt"

type Cla struct {
	n1 int
	n2 int

}

func (c Cla) sum()int{
	return c.n1 +c.n2

}

func (c Cla) mul()int{
	return c.n1 * c.n2
}

func main(){
	cc:=Cla{5,10}
	fmt.Println("sum of the given nmbr:",cc.sum())
	fmt.Println("multiplication of the given nmbr:",cc.mul())
}
