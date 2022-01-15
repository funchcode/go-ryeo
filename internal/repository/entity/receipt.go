package entity

import "time"

type Receipt struct {
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
