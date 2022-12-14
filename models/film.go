package models

import "time"

type Film struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" gorm:"type: varchar(255)"`
	CategoryID  int       `json:"category_id" form:"category_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Category    Category  `json:"category"`
	Price       int       `json:"price" gorm:"type: int"`
	Filmurl     string    `json:"filmurl" gorm:"type: varchar(255)"`
	Thumbnail   string    `json:"thumbnail" gorm:"type: varchar(255)"`
	Description string    `json:"description" gorm:"type: text"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type FilmResponse struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Category    CategoryResponse `json:"category"`
	CategoryID  int              `json:"category_id" form:"category_id"`
	Price       int              `json:"price"`
	FilmUrl     string           `json:"filmurl"`
	Description string           `json:"description"`
	Thumbnail   string           `json:"thumbnail"`
}

func (FilmResponse) TableName() string {
	return "films"
}
