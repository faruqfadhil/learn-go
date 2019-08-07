package main
import "fmt"
func swap(x int ,y string)(string,int){
	return y,x
}
func main(){
	a,b:=swap(1,"world")
	fmt.Println(a,b)
}