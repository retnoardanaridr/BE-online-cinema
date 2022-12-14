package authdto

type RegisterResponse struct {
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"type: varchar(55)"`
}

type LoginResponse struct {
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Token    string `json:"token" gorm:"type: varchar(255)"`
}

type CheckAuthResponse struct {
	Id       int    `gorm:"type: int" json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Role     string `gorm:"type: varchar(55)"  json:"role"`
}
