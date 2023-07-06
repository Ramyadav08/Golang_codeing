package main 

import "fmt"


func main(){
	ch1:=make(chan string)
	ch2:=make(chan string)
	go sending(ch1)
	getingval:= <-ch1
	fmt.Println("sending value in go func is :", getingval)
	go receiving(ch2)
	ch2 <- getingval
}

func sending(s chan string){
	s <- "ramrekha"
}

func receiving( s chan string){
	fmt.Println("receiving value is:", <- s)
}