package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res, "Welcome to our Cybersecurity Course <br>")

	fmt.Fprintf(res, "Hello Wantsome Student <br><br>")
	fmt.Fprintf(res, "<a href=\"https://wantsome.ro\"> Wantsome </a>")
}

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServeTLS(":443", "cert/server.crt", "cert/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
