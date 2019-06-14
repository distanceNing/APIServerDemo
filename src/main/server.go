package main

import (
	"../conf"
	"fmt"
	"log"
	"net/http"
	"net/rpc"

)

func testFunc() {

	// testDefer()
	// testMap()

}

// SayHello is http handle
func SayHello(w http.ResponseWriter, request *http.Request) {
	// 接受到请求 向后端转发


	//获取输入的地址是获取输入得 os 数据的 第一个位置的值
	serverAddress := conf.HttpServerConf.PortalConf.Ip + ":" +conf.HttpServerConf.PortalConf.Port
	fmt.Println("severAddress : ",serverAddress)
	// //DelayHTTP在指定的网络地址连接到HTTP RPC服务器
	///在默认HTTP RPC路径上监听。
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("发生错误了 在这里地方  DialHTTP", err)
	}
	args := ArgsTwo{1,2}
	var reply int
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Call Multiply  发生错误了哦   arith error:", err)
	}
	fmt.Printf("Arith 乘法: %d*%d=%d\n", args.A, args.B, reply)

	var quot QuotientTwo
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith 除法取整数: %d/%d=%d 余数 %d\n", args.A, args.B, quot.Quo, quot.Rem)
	w.Write([]byte("<html>\n<title>test http</title>\n Hello world\n</html>\n"))
}

func httpService(port string) {

	http.HandleFunc("/", SayHello)
	// http.ListenAndServe("")
	fmt.Printf("http server start ....\n")
	host := "127.0.0.1:"
	host += port
	fmt.Println("http will listen on ", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
type ArgsTwo struct {
	A, B int
}

type QuotientTwo struct {
	Quo, Rem int
}

func main() {


	confFileName := "/mnt/hgfs/shared/APIServerDemo/src/conf/server.conf"
	if !conf.ParseConf(confFileName) {
		log.Panicln("ParseConf fail")
		return
	}

	//httpService(conf.HttpServerConf.ListenPort)
	//testFunc()
	//httpService()
}
