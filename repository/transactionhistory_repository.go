package repository

import (
	"FinalProject4/model/entity"

	"gorm.io/gorm"
)

type TransactionHistoryRepository interface {
	CreateTransactionHistory(transactionHistory entity.TransactionHistory) (entity.TransactionHistory, error)
	GetMyTransactionHistory(idUser int) ([]entity.TransactionHistory, error)
	GetAllTransactionHistory() ([]entity.TransactionHistory, error)
}

type transactionHistoryRepository struct {
	db *gorm.DB
}

func NewTransactionHistoryRepository(db *gorm.DB) *transactionHistoryRepository {
	return &transactionHistoryRepository{db}
}

func (r *transactionHistoryRepository) CreateTransactionHistory(transactionHistory entity.TransactionHistory) (entity.TransactionHistory, error) {
	err := r.db.Create(&transactionHistory).Error
	if err != nil {
		return transactionHistory, err
	}

	return transactionHistory, nil
}

func (r *transactionHistoryRepository) GetMyTransactionHistory(idUser int) ([]entity.TransactionHistory, error) {
	var transactionHistory []entity.TransactionHistory
	err := r.db.Preload("Product.Category").Preload("User").Where("user_id = ?", idUser).Find(&transactionHistory).Error
	if err != nil {
		return transactionHistory, err
	}

	return transactionHistory, nil
}

func (r *transactionHistoryRepository) GetAllTransactionHistory() ([]entity.TransactionHistory, error) {
	var transactionHistory []entity.TransactionHistory
	err := r.db.Preload("Product.Category").Preload("User").Find(&transactionHistory).Error
	if err != nil {
		return transactionHistory, err
	}

	return transactionHistory, nil
}
