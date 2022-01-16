package repository

import (
	"Goryeo/internal/repository/scheme"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var cfg = mysql.Config{
	User:   "root",
	Passwd: "",
	Net:    "tcp",
	Addr:   "127.0.0.1:3306",
	DBName: "smoke",
}

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	} else {
		log.Println(" DB Connect ! ")
	}

	rows, err := db.Query(scheme.SchemeReceipt)
	if err != nil {
		log.Panicf(" 스키마 생성 실패 ! ")
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(" defer 에러 ")
			fmt.Println(err.Error())
		}
	}(rows)

	for rows.Next() {
		log.Println(rows.Columns())
	}
	return db
}
