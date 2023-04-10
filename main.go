package main
// import ("time")
import "fmt"

func main(){
	a := []int{90, 10, 50}
	b := a
	b[1] = 100
	fmt.Println(b)
	fmt.Println(a)
}

