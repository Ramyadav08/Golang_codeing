
package main  


import "fmt"


func main(){
	var n int =5
	IncreaseTri(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	Decrasetri(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	print123(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	NepalFlag(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	RightTriIncr(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	printRow(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	PrintCol(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	RightDecreTri(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	HalfDimTri(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	HalfRevTri(n)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	FullDiamondTri(n)
}


func IncreaseTri(n int){
	for i:=0; i<=n; i++{
		for j:=0;j<i;j++{
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}


func Decrasetri(n int){
	for i:=0;i<=n;i++{
		for j:=n; j>i;j--{
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}


func print123(n int){
	v:=1
	for i:=0;i<=n;i++{
		for j:=0;j<i;j++{
			fmt.Print(v," ")
			v++
		}
		fmt.Println("")
	}
}


func NepalFlag(n int){
	for i:=0; i<=n; i++{
		for j:=0;j<i;j++{
			fmt.Print("* ")
		}
		fmt.Println("")
	}

	for i:=1;i<=n;i++{
		for j:=n; j>i;j--{
			fmt.Print("* ")
		}
		fmt.Println("")
	}

}


func RightTriIncr(n int){
	for i:=0;i<=n;i++{
		for j:=n;j>i;j--{
			fmt.Print(" ")
		}
		for j:=0;j<i;j++{
			fmt.Print("*")
		}
		fmt.Println("")
	}
	
}


func printRow(n int){
	for i:=0;i<=n;i++{
		for j:=0;j<i;j++{
			fmt.Print(i," ")
		}
		fmt.Println("")
	}
}

func PrintCol(n int){
	for i:=0;i<=n;i++{
		for j:=1;j<=i;j++{
			fmt.Print(j," ")
		}
		fmt.Println("")
	}
}



func RightDecreTri(n int){
	for i:=0;i<=n;i++{
		for j:=0;j<i;j++{
			fmt.Print(" ")
		}
		for j:=n;j>i;j--{
			fmt.Print("*")
		}
		fmt.Println("")
	}
}



func HalfDimTri(n int){
	for i:=0;i<=n;i++{
		for j:=n;j>i;j--{
			fmt.Print(" ")
		}
		for j:=0;j<i;j++{
			fmt.Print("*")
		}
		for j:=1;j<i;j++{
			fmt.Print("*")
		}
		fmt.Println("")
	}
}


func HalfRevTri(n int){
	for i:=0;i<=n;i++{
		for j:=0;j<i;j++{
			fmt.Print(" ")
		}
		for j:=n;j>i;j--{
			fmt.Print("*")
		}
		for j:=n-1;j>i;j--{
			fmt.Print("*")
		}
		fmt.Println("")
	}

}


func FullDiamondTri(n int){
	for i:=0;i<=n;i++{
		for j:=n;j>i;j--{
			fmt.Print(" ")
		}
		for j:=0;j<i;j++{
			fmt.Print("*")
		}
		for j:=1;j<i;j++{
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i:=1;i<=n;i++{
		for j:=0;j<i;j++{
			fmt.Print(" ")
		}
		for j:=n;j>i;j--{
			fmt.Print("*")
		}
		for j:=n-1;j>i;j--{
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
