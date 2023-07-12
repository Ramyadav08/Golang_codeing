package main

import (
	"fmt"
	
	
)

func main(){
	p:= encapsulation.NewPerson("ram","yadav",23)
	
	fmt.Println("firstName:", p.GetName())
}