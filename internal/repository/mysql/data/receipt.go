package data

import (
	"Goryeo/internal/repository/entity"
	"database/sql"
	"fmt"
	"log"
	"strings"
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
	var used bool
	err := m.db.QueryRow(`SELECT * FROM RECEIPT r WHERE r.ID = ?`, receiptId).Scan(
		&id, &name, &date, &kind, &category, &contents, &amount, &assets, &used, &createdAt, &updatedAt,
	)
	if err != nil {
		log.Println(err)
	}
	log.Println(id, name, kind, category, contents, assets, date, createdAt, updatedAt, amount)
	return entity.SetAll(
		id, name, time.Now(), kind, category, contents, 0, assets, used, time.Now(), time.Now(),
	)
}

func (m MysqlReceiptData) Update(receipt entity.Receipt) {
	// todo validation check
	log.Println(" Update start! ")
	//id := receipt.GetID()
	id := "b879c7a1-5e7a-49a1-af49-d8c252608732"
	prevReceipt := m.Get(id)
	if prevReceipt == nil {
		log.Println(" 데이터가 존재하지 않습니다. ")
	}
	sb := strings.Builder{}
	sb.Grow(100)
	sb.WriteString(fmt.Sprintf("UPDATE RECEIPT SET updated_at = '%s' ", "2022-01-16"))

	if receipt.GetName() != prevReceipt.GetName() {
		sb.WriteString(fmt.Sprintf(", name = '%s' ", receipt.GetName()))
	}

	if receipt.GetDate() != prevReceipt.GetDate() {
		sb.WriteString(fmt.Sprintf(", date = '%s' ", "2022-01-16"))
	}

	if receipt.GetKind() != prevReceipt.GetKind() {
		sb.WriteString(fmt.Sprintf(", kind = '%s' ", receipt.GetKind()))
	}

	if receipt.GetCategory() != prevReceipt.GetCategory() {
		sb.WriteString(fmt.Sprintf(", category = '%s' ", receipt.GetCategory()))
	}

	if receipt.GetContents() != prevReceipt.GetContents() {
		sb.WriteString(fmt.Sprintf(", contents = '%s' ", receipt.GetContents()))
	}

	if receipt.GetAmount() != prevReceipt.GetAmount() {
		sb.WriteString(fmt.Sprintf(", amount = '%s' ", receipt.GetAmount()))
	}

	if receipt.GetAssets() != prevReceipt.GetAssets() {
		sb.WriteString(fmt.Sprintf(", assets = '%s' ", receipt.GetAssets()))
	}

	sb.WriteString(fmt.Sprintf(" where id = '%s'", id))

	_, err := m.db.Exec(sb.String())
	if err != nil {
		log.Println(" Update 중 에러가 발생했습니다. ")
		log.Println(err.Error())
	}
	log.Println(" Update end! ")
}
