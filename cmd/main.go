package main

import (
	"HelloWorldApp/internal/transport"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fmt.Println("The Server is started on 8080 port")
	mux.HandleFunc("/", transport.HelloWorld)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
