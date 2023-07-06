package main 

import "fmt"


func main(){
	var n int=5
	rectangle(n)
	fmt.Println(".......................................................................................")
	IncTri(n)
	fmt.Println(".......................................................................................")
	DrcTri(n)
	fmt.Println(".......................................................................................")
	IncDrcTri(n)
	fmt.Println(".......................................................................................")
}

func rectangle(n int){
	for i:=0;i<n;i++{
		for j:=0; j<n;j++{
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

func IncTri(n int){
	for i:=0;i<n;i++{
		for j:=0;j<i; j++{
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

func DrcTri(n int){
	for i:=0; i<n; i++{
		for j:=n; j > i ;j--{
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}


func IncDrcTri(n int){
	for i:=0;i<n;i++{
		for j:=0;j<i; j++{
			fmt.Print("* ")
		}
		fmt.Println("")
	}

	for i:=0; i<n; i++{
		for j:=n; j > i ;j--{
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}