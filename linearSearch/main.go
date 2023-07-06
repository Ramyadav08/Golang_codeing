package main

import "fmt"
func main(){
	items:=[]int{2,3,5,6,7,8,9,10,11}
	fmt.Println("the value is at the index:",linearSearch(items,8))
}

func linearSearch(datalist []int ,key int) int{
	for index,target:=range datalist{
		if target==key{
			return index
		}
	}
	return -1
}