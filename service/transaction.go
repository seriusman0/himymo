package service

import (
	"himymo/models"
	"himymo/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTransaction(transaction *models.Transaction) error {
	err := s.repo.CreateTransaction(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetTransactionById(id int) (*models.Transaction, error) {
	transaction, err := s.repo.GetTransactionById(id)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *Service) GetTransactions() ([]*models.Transaction, error) {
	transactions, err := s.repo.GetTransactions()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *Service) UpdateTransaction(transaction *models.Transaction) error {
	err := s.repo.UpdateTransaction(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTransactionById(id int) error {
	err := s.repo.DeleteTransactionById(id)
	if err != nil {
		return err
	}

	return nil
}
