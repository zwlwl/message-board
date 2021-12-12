package dao

import (
	"database/sql"
)

var dB *sql.DB

func InitDB()  {
	db,err := sql.Open("mysql","root:123456@tcp(localhost:3306)/message_board?charset=utf8mb4&parseTime=True")
	if err != nil{
		panic(err)
	}

	dB = db
}
