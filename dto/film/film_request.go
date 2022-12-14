package filmdto

type FilmRequest struct {
	Title       string `json:"title" form:"title" gorm:"type: varchar(255)"`
	CategoryID  int    `json:"category_id" form:"category_id" gorm:"type: int"`
	Price       int    `json:"price" form:"price" gorm:"type: int"`
	FilmUrl     string `json:"filmurl" form:"filmurl" gorm:"type:varchar(255)"`
	Description string `json:"description" form:"description" gorm:"type:text"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
}
