package transactiondto

import "time"

type CreateTransaction struct {
	FilmID    int       `json:"film_id" form:"film_id" gorm:"type: int"`
	Status    string    `json:"status" gorm:"type:text" form:"status"`
	Price     int       `gorm:"type: int" json:"price"`
	SellerId  int       `gorm:"type: int" json:"sellerId"`
	CreatedAt time.Time `json:"tanggal_order" form:"tanggal_order"`
}

type UpdateTransaction struct {
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status"`
	Price  int    `json:"total"`
}

type TransactionResponse struct {
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status" gorm:"type:text" form:"status"`
	FilmID int    `json:"film_id" form:"film_id" gorm:"type: int"`
}
