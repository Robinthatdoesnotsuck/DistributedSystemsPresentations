package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello retard")
	http.HandleFunc("/", say_hello)
	http.HandleFunc("/hello_other", say_hello_to_sever)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func say_hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from server uwu")
	fmt.Fprintf(w, "The resource you have requested is: %s\n", r.URL.Path)
}

func say_hello_to_sever(w http.ResponseWriter, r *http.Request) {

}
