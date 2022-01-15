package receipt

import (
	"Goryeo/internal/repository/entity"
	"fmt"
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

func (s service) Create(req CreateReceiptRequest) {
	fmt.Println(" [Receipt] start Create() ! ")
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

	fmt.Println(" [Receipt] end Create() ! ")
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
