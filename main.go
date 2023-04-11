package main
import "net/http"
// import ("time")
import "fmt"


func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello world")
	})
}

