package service

import (
	"FinalProject4/model/entity"
	"FinalProject4/model/input"
	"FinalProject4/repository"
)

type ProductService interface {
	CreateProduct(input input.ProductCreateInput) (entity.Product, error)
	FindProductByID(ID int) (entity.Product, error)
	FindProductByCategoryID(ID int) ([]entity.Product, error)
	GetProduct() ([]entity.Product, error)
	UpdateProduct(ID int, input input.ProductUpdateInput) (entity.Product, error)
	DeleteProduct(ID int) (entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) CreateProduct(input input.ProductCreateInput) (entity.Product, error) {
	product := entity.Product{
		Title:      input.Title,
		Price:      input.Price,
		Stock:      input.Stock,
		CategoryID: input.CategoryID,
	}

	newProduct, err := s.productRepository.CreateProduct(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *productService) FindProductByID(ID int) (entity.Product, error) {
	product, err := s.productRepository.FindProductByID(ID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *productService) FindProductByCategoryID(ID int) ([]entity.Product, error) {
	var product []entity.Product

	product, err := s.productRepository.FindProductByCategoryID(ID)
	if err != nil {
		return []entity.Product{}, err
	}

	return product, nil
}

func (s *productService) GetProduct() ([]entity.Product, error) {
	product, err := s.productRepository.GetProduct()

	if err != nil {
		return []entity.Product{}, err
	}

	return product, nil
}

func (s *productService) UpdateProduct(ID int, input input.ProductUpdateInput) (entity.Product, error) {
	product, err := s.productRepository.FindProductByID(ID)

	if err != nil {
		return entity.Product{}, err
	}

	updateProduct := entity.Product{
		Title:      input.Title,
		Price:      input.Price,
		Stock:      input.Stock,
		CategoryID: input.CategoryID,
	}

	product, err = s.productRepository.UpdateProduct(ID, updateProduct)

	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (s *productService) DeleteProduct(ID int) (entity.Product, error) {
	product, err := s.productRepository.DeleteProduct(ID)

	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}
