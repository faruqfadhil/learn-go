package main
import "fmt"
func split(jml int)(x,y int){
	x = jml
	y = jml - x
	return
}

func main(){
	fmt.Println(split(17))
}