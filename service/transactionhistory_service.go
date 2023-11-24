package service

import (
	"FinalProject4/model/entity"
	"FinalProject4/model/input"
	"FinalProject4/repository"
	"errors"
)

type TransactionHistoryService interface {
	CreateTransactionHistory(input input.TransactionHistoryCreateInput, idUser int, priceProduct int) (entity.TransactionHistory, error)
	GetMyTransactionHistory(idUser int) ([]entity.TransactionHistory, error)
	GetAllTransactionHistory() ([]entity.TransactionHistory, error)
}

type transactionHistoryService struct {
	transactionHistoryRepository repository.TransactionHistoryRepository
	productRepository            repository.ProductRepository
	userRepository               repository.UserRepository
}

func NewTransactionHistoryService(transactionHistoryRepository repository.TransactionHistoryRepository, productRepository repository.ProductRepository, userRepository repository.UserRepository) *transactionHistoryService {
	return &transactionHistoryService{transactionHistoryRepository, productRepository, userRepository}
}

func (s *transactionHistoryService) CreateTransactionHistory(input input.TransactionHistoryCreateInput, idUser int, priceProduct int) (entity.TransactionHistory, error) {
	productData, err := s.productRepository.FindProductByID(input.ProductID)

	userData, err := s.userRepository.FindUserByID(idUser)

	if productData.ID == 0 {
		return entity.TransactionHistory{}, errors.New("product not found")
	}

	if productData.Stock < input.Quantity {
		return entity.TransactionHistory{}, errors.New("stock is not enough")
	}

	if userData.Balance < (priceProduct * input.Quantity) {
		return entity.TransactionHistory{}, errors.New("balance is not enough")
	}

	newTransactionHistory := entity.TransactionHistory{
		ProductID:  input.ProductID,
		UserID:     idUser,
		TotalPrice: priceProduct * input.Quantity,
		Quantity:   input.Quantity,
	}

	createdTransactionHistory, err := s.transactionHistoryRepository.CreateTransactionHistory(newTransactionHistory)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	productData.Stock = productData.Stock - input.Quantity

	_, err = s.productRepository.UpdateProduct(input.ProductID, productData)

	if err != nil {
		return entity.TransactionHistory{}, err

	}

	userData.Balance = userData.Balance - (priceProduct * input.Quantity)

	_, err = s.userRepository.UpdateUser(idUser, userData)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	return createdTransactionHistory, nil
}

func (s *transactionHistoryService) GetMyTransactionHistory(idUser int) ([]entity.TransactionHistory, error) {
	transactionHistory, err := s.transactionHistoryRepository.GetMyTransactionHistory(idUser)

	if err != nil {
		return []entity.TransactionHistory{}, err
	}

	return transactionHistory, nil
}

func (s *transactionHistoryService) GetAllTransactionHistory() ([]entity.TransactionHistory, error) {
	transactionHistory, err := s.transactionHistoryRepository.GetAllTransactionHistory()

	if err != nil {
		return []entity.TransactionHistory{}, err
	}

	return transactionHistory, nil
}
