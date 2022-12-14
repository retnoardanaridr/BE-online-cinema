package categorydto

type CreateCategory struct {
	Name string `gorm:"type:varchar(255)" form:"name" json:"name"`
}

type UpdateCategoryRequest struct {
	Name string `gorm:"type:carchar(255)" form:"name" json:"name"`
}
