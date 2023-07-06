package main

import(
	"fmt"
	"encoding/json"

)
func main(){
	var wrapData  map[string]interface{}
	data:=`{"name":"ramrekha","gf":"shadi ho gai","ab axa lag raha hai":"nahi"}`
	err := json.Unmarshal([]byte(data), &wrapData)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(wrapData)

	bytesData,err:=json.Marshal(&wrapData)

	if err!=nil{
		fmt.Println("error")
	}
	fmt.Println(string(bytesData))
}