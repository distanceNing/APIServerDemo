package errno


type Error int

const (
	SUCCESS = 0
	CHANNEL_ERROR = 1000
	DB_OPERATOR_ERROR


	RPC_FAIL = 2000

)

var ErrnoMap = map[int]string{SUCCESS:"success"}
