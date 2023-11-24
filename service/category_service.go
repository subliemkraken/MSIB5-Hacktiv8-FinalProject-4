package service

import (
	"FinalProject4/model/entity"
	"FinalProject4/model/input"
	"FinalProject4/repository"
)

type CategoryService interface {
	CreateCategory(input input.CategoryCreateInput) (entity.Category, error)
	FindCategoryByID(ID int) (entity.Category, error)
	GetCategory() ([]entity.Category, error)
	UpdateCategory(ID int, input input.CategoryUpdateInput) (entity.Category, error)
	DeleteCategory(ID int) (entity.Category, error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *categoryService {
	return &categoryService{categoryRepository}
}

func (s *categoryService) CreateCategory(input input.CategoryCreateInput) (entity.Category, error) {
	newCategory := entity.Category{
		Type: input.Type,
	}

	createdCategory, err := s.categoryRepository.CreateCategory(newCategory)

	if err != nil {
		return entity.Category{}, err
	}

	return createdCategory, nil
}

func (s *categoryService) FindCategoryByID(ID int) (entity.Category, error) {
	category, err := s.categoryRepository.FindCategoryByID(ID)

	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (s *categoryService) GetCategory() ([]entity.Category, error) {
	category, err := s.categoryRepository.GetCategory()

	if err != nil {
		return []entity.Category{}, err
	}

	return category, nil
}

func (s *categoryService) UpdateCategory(ID int, input input.CategoryUpdateInput) (entity.Category, error) {
	category, err := s.categoryRepository.FindCategoryByID(ID)

	if err != nil {
		return entity.Category{}, err
	}

	updateCategory := entity.Category{
		Type: input.Type,
	}

	category, err = s.categoryRepository.UpdateCategory(ID, updateCategory)

	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (s *categoryService) DeleteCategory(ID int) (entity.Category, error) {
	category, err := s.categoryRepository.DeleteCategory(ID)

	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}
