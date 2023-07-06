// Go program to illustrate send
// and receive operation
package main

import "fmt"

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)   // we cannot declare ch <- 23 before go because the controller does not wait for the go routine so the error will be dead lock if we declare.
	ch <- 23
	fmt.Println("End Main method")
}
