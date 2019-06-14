package conf

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

// PortalConf portal conf
type PortalConf struct {
	Ip           string `xml:"ip"`
	Port         string `xml:"port"`
	Proto        string `xml:"proto"`
	ProtoModName string `xml:"proto_mod_name"`
}

// ServerConf save config
type ServerConf struct {
	ListenPort string     `xml:"listen_port"`
	RootPath   string     `xml:"root_path"`
	MysqlHost  string     `xml:"mysql_host"`
	LogPath    string     `xml:"log_path"`
	PortalConf PortalConf `xml:"portal"`
	Comment    string     `xml:",comment"`
}



// HttpServerConf cache sever conf
var HttpServerConf ServerConf

// ParseConf :parse xml config file
func ParseConf(fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("open conf file err : ", err)
		return false
	}
	fileinfo, err := file.Stat()
	if err != nil {
		log.Println("file stat err : ", err)
		return false
	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)
	file.Read(buffer)

	//fmt.Println(string(buffer))

	err = xml.Unmarshal(buffer, &HttpServerConf)
	if err != nil {
		log.Println("xml parse err : ", err)
		return false
	}

	fmt.Println(HttpServerConf)
	return true
}
