package entity

import "time"

type Receipt interface {
	GetID() string
	GetName() string
	GetDate() time.Time
	GetKind() string
	GetCategory() string
	GetContents() string
	GetAmount() rune
	GetAssets() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
type receipt struct {
	ID        string    // 식별 ID
	Name      string    // 명칭
	Date      time.Time // 날짜
	Kind      string    // 분류 todo 별도 엔터티 분리
	Category  string    // 카테고리 todo 별도 엔터티 분리
	Contents  string    // 내용
	Amount    rune      // 금액
	Assets    string    // 자산형식 todo 별도 엔터티 분리
	CreatedAt time.Time // 발행날짜
	UpdatedAt time.Time // 수정날짜
}

func (r receipt) GetID() string {
	return r.ID
}
func (r receipt) GetName() string {
	return r.Name
}
func (r receipt) GetDate() time.Time {
	return r.Date
}
func (r receipt) GetKind() string {
	return r.Kind
}
func (r receipt) GetCategory() string {
	return r.Category
}
func (r receipt) GetContents() string {
	return r.Contents
}
func (r receipt) GetAmount() rune {
	return r.Amount
}
func (r receipt) GetAssets() string {
	return r.Assets
}
func (r receipt) GetCreatedAt() time.Time {
	return r.CreatedAt
}
func (r receipt) GetUpdatedAt() time.Time {
	return r.UpdatedAt
}

func NewReceipt(name string, date time.Time, kind string, category string, contents string, amount rune, assets string) Receipt {
	return receipt{
		ID:        GenerateID(),
		Name:      name,
		Date:      date,
		Kind:      kind,
		Category:  category,
		Contents:  contents,
		Amount:    amount,
		Assets:    assets,
		CreatedAt: time.Now(),
	}
}

func SetAll(id string, name string, date time.Time, kind string, category string, contents string, amount rune, assets string, createdAt time.Time, updatedAt time.Time) Receipt {
	return receipt{
		ID:        id,
		Name:      name,
		Date:      date,
		Kind:      kind,
		Category:  category,
		Contents:  contents,
		Amount:    amount,
		Assets:    assets,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

/*
영수증
	ㄴ 식별번호: 124-1395-3938
	ㄴ 타입: 지출
	ㄴ 명칭: 장보기
	ㄴ 날짜: 2022-01-15
	ㄴ 카테고리: 쇼핑
	ㄴ 내용: 미역국 재료
	ㄴ 금액: 98500
	ㄴ 자산형식: 신용카드
	ㄴ 발행날짜: 2022-01-15
	ㄴ 수정날짜: nil
*/
