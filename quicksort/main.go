package main 

import "fmt"

func main(){
	arr:=[]int {5,4,1,2,6,3,7}
	quicksort(arr,0,(len(arr)-1))
	fmt.Println(arr)
}

func quicksort(arr[]int , lo, hi int)  {
	
	if lo>=hi {
		return
	}
	s:=lo
	e:=hi
	mid:=(s+e)/2
	p:=arr[mid]
	for s<=e {
		for arr[s]< p {
			s++
		}
		for arr[e] > p {
			e--
		}
		if s<=e {
			tep:= arr[s]
			arr[s] = arr[e] 
			arr[e] = tep
			s++
			e--
		}

	}

	// now my pivot is corresct position now two halves 
	quicksort(arr,lo,e)
	quicksort(arr,s,hi)
	
	

}