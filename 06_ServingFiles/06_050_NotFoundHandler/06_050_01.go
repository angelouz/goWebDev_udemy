package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", dog)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func dog(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL)
	fmt.Fprintln(w, "go look at your terminal")
}
