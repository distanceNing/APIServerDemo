package service

import (
	"../logger"
	"net/http"
	"net/rpc"
	"../conf"
	"github.com/gin-gonic/gin"
	"../errno"
)

var rpcClient *rpc.Client
var callPortalCount = 0


type RequestInfo struct {
	RequestIp,Args string
}

type ReplyMsg struct {
	ErrorNum    int
	ErrorInfo string
	Reply string
}

func Handler(c *gin.Context)  {

	// 接受到请求 向后端转发
	// 负载均衡
	args := &RequestInfo{c.ClientIP(),c.Request.URL.RawQuery}
	logger.Info(args)
	var reply ReplyMsg
	callBackEnd(args,&reply)
	if reply.ErrorNum != errno.SUCCESS {
		logger.Error(reply.ErrorNum,reply.ErrorInfo)
		c.JSON(http.StatusInternalServerError, gin.H{
			string(reply.ErrorNum):reply.ErrorInfo,
		})
		return
	}

	c.JSON(http.StatusOK, reply.ErrorInfo)
}


func callBackEnd(args *RequestInfo,reply *ReplyMsg){
	logger.Info("call back end service ")
	//  call back end
	serverAddress := conf.HttpServerConf.PortalConf.Ip + ":" +conf.HttpServerConf.PortalConf.Port
	logger.Debug("ctrl svr address : ",serverAddress)
	var err error
	rpcClient , err = rpc.DialHTTP("tcp", serverAddress)

	if err != nil {
		logger.Error( err)
		reply.ErrorNum = errno.RPC_FAIL
		reply.ErrorInfo = err.Error()
		return
	}

	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = rpcClient.Call("Arith.CtrlService", args, reply)


	if err != nil {
		logger.Error("Call CtrlService  error:", err)
		reply.ErrorNum = errno.RPC_FAIL
		reply.ErrorInfo = err.Error()
		return

	}
	logger.Info("Arith reply :",reply.ErrorNum,reply.ErrorInfo)
}