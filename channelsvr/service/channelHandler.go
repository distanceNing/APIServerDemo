package service

import (
	"../errno"
	"fmt"
)

type record map[string]string

func payCallback(argMap map[string]string, reply *ReplyMsg)( bool )  {

	fmt.Printf("verify order pay status")
	reply.ErrorNum = errno.SUCCESS
	reply.ErrorInfo = errno.ErrnoMap[reply.ErrorNum]
	return true
}

func payConfirm(argMap record,reply *ReplyMsg)(bool){

	return true
}


func order(argMap record,reply *ReplyMsg)(bool){
	//访问第三方渠道的支付服务
	fmt.Println("begin order")
	fmt.Println("order successful")
	return true
}

func orderHandler(argMap record, reply *ReplyMsg)( bool ){
	if !order(argMap,reply){
		fmt.Println("order fail")
		return false
	}

	// 将订单插入数据库

	return true
}

func payCancelHandler(argMap map[string]string, reply *ReplyMsg)( bool )  {
	return true
}