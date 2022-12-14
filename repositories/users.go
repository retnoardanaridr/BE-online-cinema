package repositories

import (
	"BE-waysbuck-API/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(userID int) (models.User, error)
	CreateUser(user models.User) (models.User, error) //udah ada regis
	UpdateUser(user models.User) (models.User, error) //harusnya pakai di profile
	DeleteUser(user models.User) (models.User, error) //biarin aja, :)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

// yang ini preloadnya aku masih bingung :(
// search All data
func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Transaction.Film").Preload("Transaction.User").Preload("Transaction.Film.Category").Find(&users).Error //bentar yg ini otakku ngebug

	return users, err
}

// seach data by ID
func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Transaction.Film").Preload("Transaction.User").Preload("Transaction.Film.Category").First(&user, ID).Error

	return user, err
}

// add user data
func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

// update data
func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

// delete data
func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
