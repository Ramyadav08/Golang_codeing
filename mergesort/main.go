package main 

import "fmt"


func main(){
	arr:=[]int{4,3,5,6,1,2,9,10,7}
	fmt.Println(mergesort(arr))
	// time compexity n*log(n)
}

func mergesort(array []int)[]int{
	if len(array) ==1{   // this is the base condition
		return array
	}

	mid:= len(array)/2
	left:=mergesort(array[:mid])
	right:=mergesort(array[mid:])
	return merge(left,right)
}

func merge(first,second []int)[]int{
	mix:=make([]int,len(first)+len(second))  // this is merge array after sorting
	i:=0   // this is for the left iteration
	j:=0   // this is for the right iteration
	k:=0   // this is for the mix iteration or new array

	for i<len(first) && j<len(second){  // this is the boundry loop
		if first[i] < second[j]{       // this is for the comparing values of first and second array or left and right
			mix[k]=first[i]        // if left[i]<rith[j] then assign mix or new array to i and increse i 
			i++
		}else{
			mix[k]=second[j]
			j++
		}
		k++  // this is iteration which is increment every time when the value assign
	}

	for i<len(first){   // this is for the loop if any element left in the left array
		mix[k]=first[i]
		i++
		k++
	}

	for j<len(second){   // this is for the elemnet left in the right array
		mix[k]=second[j]
		j++
		k++
	}
	return mix  // this is final output
}