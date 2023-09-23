package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type", "text/html",
	)

	bs, err := ioutil.ReadFile("hello.html")

	if err != nil {
		fmt.Println(err)
		return
	}


	io.WriteString(res, string(bs))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}