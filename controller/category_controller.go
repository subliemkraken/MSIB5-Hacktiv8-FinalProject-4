package controller

import (
	"FinalProject4/helper"
	"FinalProject4/model/input"
	"FinalProject4/model/response"
	"FinalProject4/service"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService service.CategoryService
	productService  service.ProductService
}

func NewCategoryController(categoryService service.CategoryService, productService service.ProductService) *CategoryController {
	return &CategoryController{categoryService, productService}
}

func (h *CategoryController) CreateCategory(c *gin.Context) {
	var input input.CategoryCreateInput

	err := c.ShouldBindJSON(&input)

	category, err := govalidator.ValidateStruct(input)

	if !category {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	result, err := h.categoryService.CreateCategory(input)
	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	categoryResponse := response.CategoryCreateResponse{
		ID:                result.ID,
		Type:              result.Type,
		SoldProductAmount: result.SoldProductAmount,
	}

	response := helper.APIResponse("created", categoryResponse)
	c.JSON(http.StatusCreated, response)
}

func (h *CategoryController) GetCategory(c *gin.Context) {
	category, err := h.categoryService.GetCategory()
	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	categoryResponse := []response.CategoryGetResponse{}

	for _, category := range category {
		productTemp := []response.Products{}

		product, err := h.productService.FindProductByCategoryID(category.ID)

		if err != nil {
			response := helper.APIResponse("failed", gin.H{
				"Errors": err.Error(),
			})
			c.JSON(http.StatusBadRequest, response)
			fmt.Println("error: " + err.Error())
			return
		}

		for _, product := range product {
			tmpProduct := response.Products{
				ID:        product.ID,
				Title:     product.Title,
				Price:     product.Price,
				Stock:     product.Stock,
				CreatedAt: product.CreatedAt,
				UpdatedAt: product.UpdatedAt,
			}
			productTemp = append(productTemp, tmpProduct)
		}

		tmpCategory := response.CategoryGetResponse{
			ID:                category.ID,
			Type:              category.Type,
			SoldProductAmount: category.SoldProductAmount,
			CreatedAt:         category.CreatedAt,
			UpdatedAt:         category.UpdatedAt,
			Products:          productTemp,
		}

		categoryResponse = append(categoryResponse, tmpCategory)
	}

	response := helper.APIResponse("success", categoryResponse)
	c.JSON(http.StatusOK, response)
}

func (h *CategoryController) UpdateCategory(c *gin.Context) {
	var inputCategoryUpdate input.CategoryUpdateInput

	err := c.ShouldBindJSON(&inputCategoryUpdate)

	category, err := govalidator.ValidateStruct(inputCategoryUpdate)

	if !category {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	var idCategoryUri input.CategoryUpdateID

	err = c.ShouldBindUri(&idCategoryUri)

	categoryId, err := govalidator.ValidateStruct(idCategoryUri)

	if !categoryId {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	id_category := idCategoryUri.ID

	_, err = h.categoryService.UpdateCategory(id_category, inputCategoryUpdate)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())

		return
	}

	categoryUpdated, err := h.categoryService.FindCategoryByID(id_category)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())

		return
	}

	categoryResponse := response.CategoryUpdateResponse{
		ID:                categoryUpdated.ID,
		Type:              categoryUpdated.Type,
		SoldProductAmount: categoryUpdated.SoldProductAmount,
		UpdatedAt:         categoryUpdated.UpdatedAt,
	}

	response := helper.APIResponse("success", categoryResponse)
	c.JSON(http.StatusOK, response)
}

func (h *CategoryController) DeleteCategory(c *gin.Context) {
	var idCategory input.CategoryDeleteID

	err := c.ShouldBindUri(&idCategory)

	categoryId, err := govalidator.ValidateStruct(idCategory)

	if !categoryId {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	id_category := idCategory.ID

	_, err = h.categoryService.DeleteCategory(id_category)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}
	response := helper.APIResponse("success", "Category has been successfully deleted")
	c.JSON(http.StatusOK, response)
}
