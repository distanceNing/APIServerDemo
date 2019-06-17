package service

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"../comm"
)

//成功返回 0 其他错误情况返回相应的错误码
type RequestInfo struct {
	RequestIp,Args string
}

type ReplyMsg struct {
	ErrorNum    int
	ErrorInfo string
	Reply string
}

type Arith int

func StartService() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	fmt.Println("svr will listen on 18000 ------")
	err := http.ListenAndServe(":18000", nil)

	if err != nil {
		fmt.Println("http listen ", err.Error())
	}
}

func (t *Arith) ChannelService(requestInfo *RequestInfo, reply *ReplyMsg) error {
	log.Println("call channel service")
	argMap := comm.JsonStrToMap(requestInfo.Args)
	// 向数据库插入订单信息
	cmd := argMap["cmd_code"]
	if cmd == "order"{
		if !orderHandler(argMap,reply){

		}
	}else if cmd == "pay" {
		if !payCallback(argMap,reply){

		}
	}else if cmd == "pay_cancel"{
		payCancelHandler(argMap,reply)
	}

	return nil
}
