package main

import (
	"fmt"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path !="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"method not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")

}

func formHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path !="/form"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "POST"{
		http.Error(w,"method not supported",http.StatusNotFound)
		return
	}
	if err := r.ParseForm();err != nil{
		fmt.Fprintf(w,"Parse Err: %v",err)
		return
	}
	name := r.FormValue("name")
	age := r.FormValue("age")
	fmt.Fprintf(w,"Name: %s Age: %s", name, age)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/form",formHandler)
	fmt.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil{
		log.Fatal(err)
	}
}