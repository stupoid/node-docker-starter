package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	portPtr := flag.Int("port", 80, "server port")
	flag.Parse()

	port := fmt.Sprintf(":%v", *portPtr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	fmt.Printf("Started server at http://localhost%v\n", port)
	http.ListenAndServe(port, mux)
}
