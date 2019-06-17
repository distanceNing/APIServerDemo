package main

import "../ctrlsvr"



func main() {
	/*mapResult := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	jsonStr := comm.MapToJsonStr(mapResult)


	fmt.Println(jsonStr)

	m := comm.JsonStrToMap(jsonStr)
	fmt.Println(m["France"])
	//slcD := []string{"apple", "peach", "pear"}
*/


	ctrlsvr.StartService()
}
