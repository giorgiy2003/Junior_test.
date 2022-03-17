package main

import (
	"net/http"
	
)

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8090", nil)
}

func main()
	handleRequest()
)