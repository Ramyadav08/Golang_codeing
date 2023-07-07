package main

import "fmt"

func main(){
	var n int =7
	swastikpattern(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	Xpatterb(n)

}

func swastikpattern(n int){
	mid:=n/2+1   // centre point
	for i:=1;i<=n;i++{        // for the 7 row
		for j:=1;j<=n;j++{    // for the 7 col
			if i==mid || j==mid || (i==1 && j>=mid) || (j==1 && i<=mid) || (i==n && j<=mid)||(j==n && i>=mid) {
				fmt.Print("* ")
			}else{
				fmt.Print("  ")  //2time space
			}

		}
		fmt.Println("")
	}
}



func Xpatterb(n int){
	for i:=1;i<=n;i++{
		for j:=1;j<=n;j++{
			if (i==1 && j<=n) || (i==n && j<=n) || (i==j){
				fmt.Print("* ")
			}else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}


