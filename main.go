package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "index.html")
	})
	http.ListenAndServe(":8080", nil)
	//fmt.Println("Hello World!")
}
