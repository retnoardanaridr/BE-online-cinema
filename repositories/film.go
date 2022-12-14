package repositories

import (
	"BE-waysbuck-API/models"
	"fmt"

	"gorm.io/gorm"
)

type FilmRepository interface {
	GetFilms() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	GetCategoriesbyId(CategoryID int) (models.Category, error)
	CreateFilm(film models.Film) (models.Film, error)
	UpdateFilm(film models.Film) (models.Film, error)
	DeleteFilm(film models.Film) (models.Film, error)
}

func RespositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetFilms() ([]models.Film, error) {
	var films []models.Film
	err := r.db.Preload("Category").Find(&films).Error

	return films, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
	var film models.Film
	err := r.db.Preload("Category").First(&film, ID).Error

	return film, err
}

// get for handler create film
func (r *repository) GetCategoriesbyId(CategoryID int) (models.Category, error) {
	var categories models.Category
	err := r.db.Find(&categories, CategoryID).Error

	return categories, err
}

func (r *repository) CreateFilm(film models.Film) (models.Film, error) {
	err := r.db.Create(&film).Error
	fmt.Println(film)

	return film, err
}

func (r *repository) UpdateFilm(film models.Film) (models.Film, error) {
	r.db.Model(&film).Association("Category").Replace(film.Category)
	err := r.db.Save(&film).Error

	return film, err
}

func (r *repository) DeleteFilm(film models.Film) (models.Film, error) {
	err := r.db.Delete(&film).Error

	return film, err
}
