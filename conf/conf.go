package conf

import (
	"log"
	"os"
)

// ServerConf save config
type ServerConf struct {
	listenPort int
	rootName   string
}

var serverConf ServerConf

// read json config file
func readConf(fileName string) bool {
	f, err := os.OpenFile(fileName, 1, 0)
	if err != nil {
		log.Println("open file fail err is : ", err)
	}
	var buffer [1024]byte
	n, err := f.Read(buffer)

	return true
}
