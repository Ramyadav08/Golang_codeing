package main

import "fmt"

func main(){
	arr:=[]int{7,5,6,2,1,3,4}
	cyclesort(arr)
	fmt.Println(arr)
}

func cyclesort(arr []int){
	i:=0
	for i<len(arr) {
		correctIndex:=arr[i]-1
		if arr[i] !=arr[correctIndex]{
			swap(arr,i,correctIndex)
		}else {
			i++
		}

	}
}

func swap(arr []int,i,correctIndex int){
	temp:=arr[i]
	arr[i]=arr[correctIndex]
	arr[correctIndex]=temp
}

// time complexity n^2
