package receipt

import (
	"Goryeo/internal/entity"
	"fmt"
	"time"
)

type Receipt struct {
	entity.Receipt
}

type CreateReceiptRequest struct {
	Name     string
	Date     time.Time
	Kind     string
	Category string
	Contents string
	Amount   rune
	Assets   string
}

func Create(req CreateReceiptRequest) {
	/*
		todo Request validation 체크
		if err = req.validate ; err != nil {
			return err
		}
	*/
	ID := entity.GenerateID()
	receipt := entity.Receipt{
		ID:        ID,
		Name:      req.Name,
		Kind:      req.Kind,
		Date:      req.Date,
		Category:  req.Category,
		Contents:  req.Contents,
		Amount:    req.Amount,
		Assets:    req.Assets,
		CreatedAt: time.Now(),
	}
	fmt.Println(receipt)
}
