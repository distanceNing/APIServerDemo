package comm



import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
// 创建一个协程去做数据库操作
func Connectdb()  {
	db, err := sql.Open("mysql", "root:nimapi@tcp(127.0.0.1:3306)/temp?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	sqlResult,err := db.Exec("select * from stu;")

	if err != nil {
		log.Println("access db fail ",err.Error())
	}
	fmt.Println(sqlResult)
}

