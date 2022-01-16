package data

import (
	"Goryeo/internal/repository/entity"
	"database/sql"
	"log"
	"time"
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

func (m MysqlReceiptData) Get(receiptId string) entity.Receipt {
	log.Println(receiptId)
	//rows, err := m.db.Query(`SELECT * FROM RECEIPT r WHERE r.ID = $1`, receiptId)
	var id, name, kind, category, contents, assets string
	var date, createdAt string
	var updatedAt sql.NullString
	var amount int64
	err := m.db.QueryRow(`SELECT * FROM RECEIPT r WHERE r.ID = ?`, receiptId).Scan(
		&id, &name, &date, &kind, &category, &contents, &amount, &assets, &createdAt, &updatedAt,
	)
	if err != nil {
		log.Println(err)
	}
	log.Println(id, name, kind, category, contents, assets, date, createdAt, updatedAt, amount)
	return entity.SetAll(id, name, time.Now(), kind, category, contents, 0, assets, time.Now(), time.Now())
}
