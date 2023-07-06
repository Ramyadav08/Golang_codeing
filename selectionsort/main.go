package main 

import "fmt"

func main(){
	arr:=[]int{4,2,7,1,9}
	fmt.Println(SelectionSort(arr))
}


func SelectionSort(array[] int)[]int{
	for i:=0;i<len(array)-1;i++{
		min:=i
		for j:=i+1; j<len(array);j++{
			if array[j]<array[min]{
				min=j
			}
		}

		temp:=array[min]
		array[min]=array[i]
		array[i]=temp
	}

	return array
}