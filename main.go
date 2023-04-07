package main
import "fmt"

//pointers
func main(){
	var color string = "Green"
	fmt.Println(color)
	pointerpointer(&color)
	fmt.Println(color)
}

func pointerpointer(s *string){
	newvalue := "Red"
	*s = newvalue
}