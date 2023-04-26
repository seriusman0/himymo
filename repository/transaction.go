package repository

import (
	"fmt"

	"gorm.io/gorm"
	"himymo/models"
)

type Repository struct {
	DB *gorm.DB
}

func (repo *Repository) Migrate() error {
	return repo.DB.AutoMigrate(&models.Transaction{})
}

func (repo *Repository) CreateTransaction(transaction *models.Transaction) error {
	return repo.DB.Create(transaction).Error
}

func (repo *Repository) GetTransactionById(id int) (*models.Transaction, error) {
	var transaction models.Transaction

	err := repo.DB.First(&transaction, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("transaction not found with ID %d", id)
		}

		return nil, fmt.Errorf("failed to get transaction with ID %d: %v", id, err)
	}

	return &transaction, nil
}

func (repo *Repository) GetTransactions() ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	err := repo.DB.Find(&transactions).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %v", err)
	}

	return transactions, nil
}

func (repo *Repository) UpdateTransaction(transaction *models.Transaction) error {
	return repo.DB.Save(transaction).Error
}

func (repo *Repository) DeleteTransactionById(id int) error {
	return repo.DB.Delete(&models.Transaction{}, id).Error
}
