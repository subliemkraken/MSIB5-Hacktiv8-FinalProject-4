package repository

import (
	"FinalProject4/model/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product entity.Product) (entity.Product, error)
	FindProductByID(ID int) (entity.Product, error)
	FindProductByCategoryID(ID int) ([]entity.Product, error)
	GetProduct() ([]entity.Product, error)
	UpdateProduct(ID int, product entity.Product) (entity.Product, error)
	DeleteProduct(ID int) (entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) CreateProduct(product entity.Product) (entity.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) FindProductByID(ID int) (entity.Product, error) {
	var product entity.Product
	err := r.db.Preload("Category").Where("id = ?", ID).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) GetProduct() ([]entity.Product, error) {
	var product []entity.Product
	err := r.db.Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) FindProductByCategoryID(ID int) ([]entity.Product, error) {
	var product []entity.Product

	err := r.db.Preload("Category").Where("category_id = ?", ID).Find(&product).Error

	if err != nil {
		return []entity.Product{}, err
	}

	return product, nil
}

func (r *productRepository) UpdateProduct(ID int, product entity.Product) (entity.Product, error) {
	err := r.db.Where("id = ?", ID).Updates(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) DeleteProduct(ID int) (entity.Product, error) {
	var product entity.Product
	err := r.db.Where("id = ?", ID).Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
