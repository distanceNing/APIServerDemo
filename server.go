package main

import (
	"fmt"
	"log"
	"net/http"
)

func testFunc() {

	testDefer()
	// testMap()

}

// SayHello is http handle
func SayHello(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("<html><title>test http</title> Hello world</html>"))
}

func main() {
	testFunc()

	http.HandleFunc("/", SayHello)
	// http.ListenAndServe("")
	fmt.Printf("http server start ....\n")
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fmt.Printf("http listen on 8000")
}
