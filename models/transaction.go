package models

import "time"

type Transaction struct {
	ID        int64        `json:"id"` //aku add 64 nya
	UserID    int          `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      User         `json:"user"`
	Status    string       `json:"status"`
	Price     int          `json:"price"`
	FilmID    int          `json:"film_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Film      FilmResponse `json:"film"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}

type TransactionResponse struct {
	ID     int          `json:"id"`
	UserID int          `json:"user_id"`
	User   User         `json:"user"`
	FilmID int          `json:"film_id"`
	Film   FilmResponse `json:"film"`
	Price  int          `json:"price"`
	Status string       `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
