package main
import "fmt"
func main(){
	jml:=1
	for ; jml < 10 ; {
		jml+=jml
	}
	fmt.Println(jml)
}