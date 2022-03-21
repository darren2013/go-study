package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {

	flag.Set("v", "4")

	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func rootHandler(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	fmt.Println("root handler....")
	user := httpRequest.URL.Query().Get("u")

	if user != "" {
		io.WriteString(responseWriter, fmt.Sprintf("hello %s \n", user))
	} else {
		io.WriteString(responseWriter, "hello,[stranger] \n")
	}

	io.WriteString(responseWriter, "details of the request header")

	for k, v := range httpRequest.Header {
		io.WriteString(responseWriter, fmt.Sprintf("%s=%s", k, v))
	}

}
