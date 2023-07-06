package main

import "fmt"
func main(){
	items:=[]int{2,3,4,5,6,7,8,9,11,12,13,14}
	fmt.Println("the value is at index:",birearySearch(items,13))
}

func birearySearch(datalist []int,target int)int{
	var start int=0
	var last int=len(datalist)-1
	for start<=last {
		var mid int=start + (last-start)/2;
		if target < datalist[mid]{
			last=mid -1
		}else if target > datalist[mid] {
			start=mid+1
		} else {
			return mid
		}
		
	}
	return -1
}