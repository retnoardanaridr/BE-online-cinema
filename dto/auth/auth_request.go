package authdto

type RegisterRequest struct {
	Fullname string `json:"fullname" gorm:"type: varchar(255) required"`
	Email    string `json:"email" gorm:"type: varchar(255) required"`
	Password string `json:"password" gorm:"type: varchar(255) required"`
}

type LoginRequest struct {
	Email    string
	Password string
}
