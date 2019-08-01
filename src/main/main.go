package main

import (
	"../router"
	"../logger"
	"../conf"
)

func main() {

	confFileName := "/home/yn/APIServerDemo/src/conf/server.conf"
	if !conf.ParseConf(confFileName) {
		logger.Error("ParseConf fail")
		return
	}


	host := ":" + conf.HttpServerConf.ListenPort

	router.InitRouter().Run(host) // listen and serve on 0.0.0.0:8080
}
