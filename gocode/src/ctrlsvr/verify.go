package ctrlsvr

import (
	//"strings"
)

func VerifyRequest(argMap map[string]string ,reply *ReplyMsg)( bool)  {
	//检查渠道的合法性

	if argMap["pay_channel"] != "wechatpay"{
		reply.ErrorInfo = "request channel is invaild"
		return false
	}


	return true
}