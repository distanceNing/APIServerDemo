package ctrlsvr

import (
	"log"
	"net/rpc"
	"../comm"
	"../errno"
)

func AccessRiskCtrl(argMap map[string]string) (bool) {
	log.Println("access risk ctrl svr")
	return true
}




func channelHandler(argMap map[string]string, reply *ReplyMsg) (bool) {
	rpcClient , err := rpc.DialHTTP("tcp", "127.0.0.1:18000")
	if err != nil {
		log.Println("DialHTTP : ", err)
		reply.ErrorNum = errno.RPC_FAIL
		reply.ErrorInfo = err.Error()
		return false
	}
	requestInfo := RequestInfo{"127.0.0.1",comm.MapToJsonStr(argMap)}

	rpcClient.Call("Arith.ChannelService",requestInfo,reply)

	if err != nil {
		log.Fatal("Call ChannelService error:", err)
	}
	log.Printf("Channle svr reply :",reply)

	return true
}



func queryOrderHandler(argMap map[string]string, reply *ReplyMsg)(bool){
	return true
}

func otherServiceHandler(argMap map[string]string, reply *ReplyMsg) (bool) {
	reply.ErrorNum = 1
	reply.ErrorInfo = "success"
	return true
}