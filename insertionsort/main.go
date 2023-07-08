package main

import "fmt"


func main(){
	arr:=[]int {5,3,4,1,2}
	insertionsort(arr)
	//fmt.Println(arr)
}

func insertionsort(arr []int){
	for i:=1;i<len(arr);i++ {
		for j:=0; j < i ; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i],arr[j]

			}

		}
	}
	fmt.Println(arr)
}