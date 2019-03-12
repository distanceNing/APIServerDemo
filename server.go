package main

import (
	"fmt"
	"log"
	"net/http"
)

// SayHello is http handle
func SayHello(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("<html><title>test http</title> Hello world</html>"))
}

func main() {
	http.HandleFunc("/", SayHello)
	// http.ListenAndServe("")
	fmt.Printf("http server start ....\n")
	if err := http.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fmt.Printf("http listen on 8000")
}
