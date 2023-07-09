package main 

import "fmt"

func main(){
	arr:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19}
	targedt:= 10
	fmt.Println("the element is at index:",infiniteArr(arr,targedt))
}

func infiniteArr(arr []int, target int)int{
	//first find the range 
	// first craete the box size
	s:=0
	e:=1

	for target > arr[e] {
		news:=e+1
		e= e+(e-s+1)*2 // doubling the size of the array
		s=news
	}
	return binearySearch(arr, target,s,e)
		
}


func binearySearch(arr []int,tar,s,e int)int{
	mid:=s+(e-s)/2
	for s<=e {
		if tar < arr[mid]{
			e--
		}else if tar > arr[s]{
			s++
		}else {
			return mid
		}
	}
	return -1
}