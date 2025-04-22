package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("hello world\n")
	serveMux := http.NewServeMux()
	var this_server http.Server
	this_server.Handler = serveMux
	// serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	this_server.Addr = ":8080"

	this_server.ListenAndServe()
	for true {
	}

}
