package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the sum of 2 + 2 = %d", sum))
}

func addValues(x int, y int) int {
	return x + y
}

func main(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
