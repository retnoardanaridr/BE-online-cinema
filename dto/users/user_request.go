package usersdto

type CreateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"id"`
}

type UpdateUserRequest struct {
	Fullname string `json:"fullname" gorm:"type: varchar(255)" validate:"required"`
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
	Image    string `json:"Image" gorm:"type: varchar(255)" validate:"required"`
}
