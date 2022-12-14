package repositories

import (
	"BE-waysbuck-API/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)             //all trans
	GetTransactionByUser(userID int) (models.Transaction, error) //buat yg udah beli
	// GetTransaction(ID int) (models.Transaction, error)
	GetTransaction() (models.Transaction, error) //midtrans
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransactionID(ID int) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error) // midtrans notif  Declare GetOneTransaction repository method
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransactions(status string, ID string) error //midtrans notifications
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Film.Category").Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransactionByUser(userID int) (models.Transaction, error) {
	var transactions models.Transaction
	err := r.db.Preload("Film.Category").Preload("User").Where("user_id =?", userID).First(&transactions).Error

	return transactions, err
}

// func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
// 	var transactions models.Transaction
// 	err := r.db.Preload("Film").First(&transactions, "id = ?", ID).Error

// 	return transactions, err
// }

func (r *repository) GetTransactionID(ID int) (models.Transaction, error) {
	var transactionId models.Transaction

	err := r.db.Preload("Film").First(&transactionId, ID).Error
	return transactionId, err
}

func (r *repository) GetTransaction() (models.Transaction, error) {
	var transactions models.Transaction
	err := r.db.Preload("User").Preload("Film").Find(&transactions, "status = ?", "waiting").Error

	return transactions, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransactions(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("Film").First(&transaction, ID)

	// If is different & Status is "success" decrement film quantity
	if status != transaction.Status && status == "success" {
		var film models.Film
		r.db.First(&film, transaction.ID)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}

// GetOneTransaction method here ...
func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Film").Preload("Film.User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}
