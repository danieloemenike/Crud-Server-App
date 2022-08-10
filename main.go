package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/danieloemenike/router"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static/sprWeb"))
	http.Handle("/",fileServer)
	http.HandleFunc("/hello",helloHandler)
	fmt.Printf("Starting server at port 8080\n")
	r:= router.Router()
	fmt.Println("Starting the server on port 9000")
	if err:=http.ListenAndServe(":9000",r);err!=nil{
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter,r *http.Request){
	//if the path is not hello, then we should display an error message
	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"Method is not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"hello! i m here")
}