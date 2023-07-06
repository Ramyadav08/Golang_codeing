
// package main

// import "fmt"

// func display(str string) {
// 	for w := 0; w < 6; w++ {
// 		fmt.Println(str)
// 	}
// }

// func main() {

// 	// Calling Goroutine
// 	go display("Welcome")

// 	// Calling normal function
// 	display("GeeksforGeeks")
// 	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
// }

/*In the above example, we simply create a display() function and then call this function in two different ways first one is a Goroutine, i.e. go display(“Welcome”) and another one is a normal function, i.e. display(“GeeksforGeeks”). But there is a problem, it only displays the result of the normal function that does not display the result of Goroutine because when a new Goroutine executed, the Goroutine call return immediately. The control does not wait for Goroutine to complete their execution just like normal function they always move forward to the next line after the Goroutine call and ignores the value returned by the Goroutine. So, to executes a Goroutine properly, we made some changes in our program as shown in the below code:*/


// Go program to illustrate the concept of Goroutine
package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main() {

	// Calling Goroutine
	go display("Welcome")

	// Calling normal function
	display("GeeksforGeeks")
}
