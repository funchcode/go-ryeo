package receipt

import (
	"Goryeo/internal/repository/entity"
	"Goryeo/internal/repository/mysql/data"
	"fmt"
	"log"
	"time"
)

type Service interface {
	Create(req CreateReceiptRequest)
	Get(id string)
	Query()
	Update()
	Delete()
}

type service struct {
}

func NewService() Service {
	return service{}
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

func (s service) Create(req CreateReceiptRequest) {
	log.Println(" [Receipt] start Create() ! ")
	/*
		todo Request validation 체크
		if err = req.validate ; err != nil {
			return err
		}
	*/
	receipt := entity.NewReceipt(req.Name, req.Date, req.Kind, req.Category, req.Contents, req.Amount, req.Assets)
	data.NewReceiptData()
	log.Println(receipt)

	log.Println(" [Receipt] end Create() ! ")
}

func (s service) Get(id string) {
	fmt.Println(" [Receipt] start Get() ! ")

	fmt.Println(" [Receipt] end Get() ! ")
}

func (s service) Query() {

}

func (s service) Update() {

}

func (s service) Delete() {

}
