package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func getRoot(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is my website")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":8090", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

// mux := http.NewServeMux()
// mux.HandleFunc("/", getRoot)
// mux.HandleFunc("/hello", getHello)
//
// err := http.ListenAndServe(":3333", mux)
