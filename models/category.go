package models

import "time"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `jsons:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `jsons:"name"`
}

func (CategoryResponse) TableName() string {
	return "category"
}
