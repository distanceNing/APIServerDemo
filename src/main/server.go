package main

import (
	"conf"
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

func httpService(port string) {

	http.HandleFunc("/", SayHello)
	// http.ListenAndServe("")
	fmt.Printf("http server start ....\n")
	host := "127.0.0.1:"
	host += port
	fmt.Println("http wile listen on ", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func main() {
	confFileName := "../conf/server.conf"
	if !conf.ParseConf(confFileName) {
		log.Panicln("ParseConf fail")
		return
	}

	httpService(conf.HttpServerConf.ListenPort)
	//testFunc()
	//httpService()
}
