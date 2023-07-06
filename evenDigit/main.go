package main
import "fmt"
func main(){
	items:=[]int{12,3,9,123,1234,2323}
	fmt.Println(findNumber(items))
}

func findNumber(datalist []int)int{
	count:=0
	for _,n:=range datalist{
		if even(n) {
			count ++
		}
	}
	return count 
}

// check wether num contain even nmbr or not
func even(n int)bool{
	var nmbrOfDigit int =digit(n)
	return nmbrOfDigit % 2==0
}

func digit(n int)int{
	count:=0
	for {
		n=n/10
		count ++
		if n==0{
			break
		}
	}
	return count
}