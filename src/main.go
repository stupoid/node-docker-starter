package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	port := ":8000"
	if len(os.Args) > 1 {
		port = fmt.Sprintf(":%v", os.Args[1])
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	fmt.Printf("Started server at http://localhost%v.\n", port)
	http.ListenAndServe(port, mux)
}
