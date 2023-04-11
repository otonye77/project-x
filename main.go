package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

func main(){
	http.HandleFunc("/", Home)
	_ = http.ListenAndServe(":8080", nil)
}

