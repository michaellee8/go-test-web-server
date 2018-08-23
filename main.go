package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, dbErr := sql.Open("mysql", "ml8_admin:ml8-mysql-1026@(db4free.net:3306)/ml8_mysql_server")
	if dbErr != nil {
		panic(dbErr)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	fmt.Println("db sucessfully connected")
	
}
