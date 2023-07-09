package main 

import "fmt"

func main(){
	var n int=1
	print1(n)
	fmt.Println("this is the fibbonachhi series:",fibbo(11))
	fmt.Println("the factorial of the given number is:", factorial(5))
}

func print1(n int){
	if n==5{              // this is base condition to stop the function 
		fmt.Println(5)
		return
	}

	fmt.Println(n)  // this is used to print the number
	print1(n+1)  // this is recursive function 
}

func fibbo(n int)int{
	if n<2{
		return n
	}
	return fibbo(n-1)+fibbo(n-2)


}

func factorial(n int)int{
	if n<=1{
		return 1
	}
	return n*factorial(n-1)
}