package data

import (
	"Goryeo/internal/repository/entity"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"time"
)

func TestNewReceiptData(t *testing.T) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "smoke",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf(" Connection Error. ")
	}
	receipt := entity.NewReceipt("장보기", time.Now(), "지출", "쇼핑", "", 98000, "신용카드")
	println(receipt.GetCategory())
	println(receipt.GetName())
	println(receipt.GetKind())
	NewReceiptData(db).Save(receipt)
}
