package data

import (
	"Goryeo/internal/repository/entity"
	"database/sql"
	"log"
)

func NewReceiptData(db *sql.DB) *MysqlReceiptData {
	return &MysqlReceiptData{
		db: db,
	}
}

type MysqlReceiptData struct {
	db *sql.DB
}

func (m MysqlReceiptData) Save(receipt entity.Receipt) {
	log.Println(receipt)
	_, err := m.db.Exec(`INSERT INTO RECEIPT 
    	(id, name, date, kind, category, contents, amount, assets, created_at, updated_at) 
    	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		receipt.GetID(), receipt.GetName(), receipt.GetDate(), receipt.GetKind(), receipt.GetCategory(), receipt.GetContents(), receipt.GetAmount(), receipt.GetAssets(), receipt.GetCreatedAt(), nil)
	if err != nil {
		log.Fatalln(err)
		log.Fatalln(" INSERT FAIL. ")
	}
}
