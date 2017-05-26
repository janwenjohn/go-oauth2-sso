package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"../util"
)

var Conn *sql.DB

//TODO panic when conn is nil
func init() {
	conn, err := sql.Open("mysql", util.DB.Db_user + ":" + util.DB.Db_password+
		"@tcp("+ util.DB.Db_host+ ":"+ util.DB.Db_port+ ")/"+ util.DB.Db_database+ "?charset=utf8")
	if err != nil {
		fmt.Println("error while connect to mysql:")
		fmt.Println(err)
	}
	conn.SetMaxOpenConns(util.DB.Db_max_open)
	conn.SetMaxIdleConns(util.DB.Db_max_idle)
	fmt.Println("conn created:")
	fmt.Println(conn)
	Conn = conn
}
