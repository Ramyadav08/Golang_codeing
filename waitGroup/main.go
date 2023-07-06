package main 


import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
func main(){
	website:=[]string{
		"https://google.com",
		"https://github.com",
	}

	for _,web := range website {
		go getStatuseCode(web)
		wg.Add(1)
		
	}
	wg.Wait()
}

func getStatuseCode(endpoint string){
	defer wg.Done()
	res,err:=http.Get(endpoint)

	if err!=nil{
		fmt.Println("Ooops in endpoint")
	}else{
		fmt.Printf("%d status code is %s \n",res.StatusCode,endpoint)
	}
}