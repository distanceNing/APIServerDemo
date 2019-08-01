package ctrlsvr

import "log"

type State int

// iota 初始化后会自动递增
const (
	kInitState State = iota
	kDispatch
	kCallChannel
	kQueryOrder
	kOtherService
	kExitState
	kDelivery
)

func Dispatch(argMap map[string]string, reply *ReplyMsg) (bool) {
	log.Println("dispatch")
	curState := kInitState
	for ; curState != kExitState; {
		switch curState {
		case kInitState:
			//风控
			log.Println("begin access risk ctr svr")
			if !AccessRiskCtrl(argMap) {
				log.Println("access risk ctl svr error")
				curState = kExitState
			}
			curState = kDispatch
		case kDispatch :
			log.Println("cmd_code is : ",argMap["cmd_code"])

			if argMap["cmd_code"] == "pay" || argMap["cmd_code"] == "order" {
				curState = kCallChannel
			}else if argMap["cmd_code"] == "query"{
				curState = kQueryOrder
			}else {
				curState = kOtherService
			}

		case kCallChannel: //调用渠道
			if !channelHandler(argMap,reply){
				log.Println("call channel fail",reply.ErrorInfo)
			}

			if argMap["cmd_code"] == "pay"{
				curState = kDelivery
			}
			curState = kExitState
		case kQueryOrder:
			queryOrderHandler(argMap,reply)
		case kOtherService:
			otherServiceHandler(argMap,reply)
			log.Println("other request service")
			curState = kExitState
		case kDelivery:
			if !deliveryHandler(argMap,reply){
				log.Println("call channel fail",reply.ErrorInfo)
			}
			curState = kExitState
		default:
			log.Println("unknow request service")
		}
	}

	return true
}

