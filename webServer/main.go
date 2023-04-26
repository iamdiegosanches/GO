package main

import (
	"fmt"
	"log"
	"net/http"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	errorCheck(err)
	fmt.Fprintf(writer, "POST request successful\n")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(writer, "Name = %s\n", name)
	fmt.Fprintf(writer, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
