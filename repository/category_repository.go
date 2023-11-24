package repository

import (
	"FinalProject4/model/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category entity.Category) (entity.Category, error)
	FindCategoryByID(ID int) (entity.Category, error)
	GetCategory() ([]entity.Category, error)
	UpdateCategory(ID int, category entity.Category) (entity.Category, error)
	DeleteCategory(ID int) (entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) FindCategoryByID(ID int) (entity.Category, error) {
	var category entity.Category
	err := r.db.Where("id = ?", ID).First(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) GetCategory() ([]entity.Category, error) {
	var category []entity.Category
	err := r.db.Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) UpdateCategory(ID int, category entity.Category) (entity.Category, error) {
	err := r.db.Where("id = ?", ID).Updates(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) DeleteCategory(ID int) (entity.Category, error) {
	var category entity.Category
	err := r.db.Where("id = ?", ID).Delete(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
