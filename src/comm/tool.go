package comm


import (
	"strings"
	"encoding/json"
	"../logger"
	"log"
)


// user=Summer&passwd=12345678
func StrToMap(str string)(map[string]string)  {
	var argMap map[string]string
	argMap = make(map[string]string)
	for _, v := range strings.Split(str, "&") {
		strs := strings.Split(v,"=")
		argMap[strs[0]] = strs[1]
	}

	return  argMap
}

func JsonStrToMap(jsonStr string) (map[string]string) {
	var f interface{}
	var data []byte = []byte(jsonStr)
	err := json.Unmarshal(data, &f)
	if err != nil {
		logger.Error(err)
	}

	m := f.(map[string]interface{})
	res := make(map[string]string)
	for k, v := range m {
		res[k] = v.(string)
	}
	return res
}

func MapToJsonStr(kv map[string]string) (string){
	jsonStr, err := json.Marshal(kv)
	if err != nil{
		logger.Error("json encode error",err.Error())
		return ""
	}
	return string(jsonStr)
}
