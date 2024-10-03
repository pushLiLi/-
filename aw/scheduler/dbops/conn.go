package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err    error
)

func init() {
	//初始化连接池，整个dbpos中都可以使用
	dbConn, err = sql.Open("mysql",
		"root:root12346@tcp(116.196.80.172:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err)
	}
}
