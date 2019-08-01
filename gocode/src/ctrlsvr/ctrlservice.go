package ctrlsvr


import (
	"fmt"
	"net/rpc"
	"net/http"
	"log"
	"../comm"
)

//成功返回 0 其他错误情况返回相应的错误码
type ReplyMsg struct {
	ErrorNum  int
	ErrorInfo string
	Reply string
}

type RequestInfo struct {
	RequestIp string
	Args string
}

type Arith string

func StartService() {
	arith:=new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	fmt.Println("svr will listen on 8000 ------")
	err:=http.ListenAndServe(":8081",nil)

	if err != nil {
		fmt.Println("http listen ",err.Error())
	}
}

func (t *Arith) CtrlService(requestInfo *RequestInfo, reply *ReplyMsg) error {
	log.Println("call ctrl service")
	//业务校验模块
	argMap := comm.StrToMap(requestInfo.Args)
	if !VerifyRequest(argMap,reply){
		log.Println("[error]",reply.ErrorInfo)
		return nil
	}
	//业务分发模块 根据请求判断请求的业务功能.并且转发到业务交互模块
	if !Dispatch(argMap,reply){
		log.Println("[error]",reply.ErrorInfo)
		return nil
	}
	return nil
}

