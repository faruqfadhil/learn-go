package main

import "fmt"
func loop(jml int) int{
	for i:=0;i < 10; i++{
		jml+=i
	}
	return jml	
}

func main(){
	fmt.Println(loop(0))
}