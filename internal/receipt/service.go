package receipt

import (
	"Goryeo/internal/repository/entity"
	"Goryeo/internal/repository/mysql/data"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Service interface {
	Create(req CreateReceiptRequest)
	Get(id string) Response
	Query()
	Update()
	Delete()
}

type service struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &service{
		db: db,
	}
}

type CreateReceiptRequest struct {
	Name string `json:"name"`
	//Date     time.Time `json:"date"`
	Kind     string `json:"kind"`
	Category string `json:"category"`
	Contents string `json:"contents"`
	Amount   int    `json:"amount"`
	Assets   string `json:"assets"`
}

type Response struct {
	ID string
	CreateReceiptRequest
}

func (s service) Create(req CreateReceiptRequest) {
	log.Println(" [Receipt] start Create() ! ")
	/*
		todo Request validation 체크
		if err = req.validate ; err != nil {
			return err
		}
	*/
	//receipt := entity.NewReceipt(req.Name, req.Date, req.Kind, req.Category, req.Contents, req.Amount, req.Assets)
	receipt := entity.NewReceipt(req.Name, time.Now(), req.Kind, req.Category, req.Contents, 0, req.Assets)
	data.NewReceiptData(s.db).Save(receipt)
	log.Println(receipt)

	log.Println(" [Receipt] end Create() ! ")
}

func (s service) Get(id string) Response {
	fmt.Println(" [Receipt] start Get() ! ")
	receipt := data.NewReceiptData(s.db).Get(id)
	fmt.Println(" [Receipt] end Get() ! ")
	return Response{
		receipt.GetID(),
		CreateReceiptRequest{
			"",
			"",
			"",
			"",
			0,
			"",
		},
	}
}

func (s service) Query() {

}

func (s service) Update() {

}

func (s service) Delete() {

}
