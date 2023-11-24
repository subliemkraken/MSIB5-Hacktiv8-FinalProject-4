package controller

import (
	"FinalProject4/helper"
	"FinalProject4/model/entity"
	"FinalProject4/model/input"
	"FinalProject4/model/response"
	"FinalProject4/service"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService}
}

func (h *ProductController) CreateProduct(c *gin.Context) {
	var input input.ProductCreateInput

	err := c.ShouldBindJSON(&input)

	product, err := govalidator.ValidateStruct(input)

	if !product {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	result, err := h.productService.CreateProduct(input)
	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	productResponse := response.ProductCreateResponse{
		ID:         result.ID,
		Title:      result.Title,
		Price:      result.Price,
		Stock:      result.Stock,
		CategoryID: result.CategoryID,
		CreatedAt:  result.CreatedAt,
	}

	response := helper.APIResponse("created", productResponse)
	c.JSON(http.StatusCreated, response)
}

func (c *ProductController) FindProductByID(id int) (entity.Product, error) {
	product, err := c.productService.FindProductByID(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (h *ProductController) GetProduct(c *gin.Context) {
	product, err := h.productService.GetProduct()
	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	productResponse := []response.ProductGetResponse{}

	for _, product := range product {
		productResponseTemp := response.ProductGetResponse{
			ID:         product.ID,
			Title:      product.Title,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryID: product.CategoryID,
			CreatedAt:  product.CreatedAt,
		}

		productResponse = append(productResponse, productResponseTemp)
	}

	response := helper.APIResponse("success", productResponse)
	c.JSON(http.StatusOK, response)
}

func (h *ProductController) UpdateProduct(c *gin.Context) {
	var inputProductUpdate input.ProductUpdateInput

	err := c.ShouldBindJSON(&inputProductUpdate)

	product, err := govalidator.ValidateStruct(inputProductUpdate)

	if !product {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	var idProductUri input.ProductUpdateID

	err = c.ShouldBindUri(&idProductUri)

	productId, err := govalidator.ValidateStruct(idProductUri)

	if !productId {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	id_product := idProductUri.ID

	_, err = h.productService.UpdateProduct(id_product, inputProductUpdate)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	productUpdated, err := h.productService.FindProductByID(id_product)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())

		return
	}

	productResponse := response.ProductUpdateResponse{
		ID:         productUpdated.ID,
		Title:      productUpdated.Title,
		Price:      productUpdated.Price,
		Stock:      productUpdated.Stock,
		CategoryID: productUpdated.CategoryID,
		CreatedAt:  productUpdated.CreatedAt,
		UpdatedAt:  productUpdated.UpdatedAt,
	}

	response := helper.APIResponse("success", productResponse)
	c.JSON(http.StatusOK, response)
}

func (h *ProductController) DeleteProduct(c *gin.Context) {
	var idProduct input.ProductDeleteInput

	err := c.ShouldBindUri(&idProduct)

	categoryId, err := govalidator.ValidateStruct(idProduct)

	if !categoryId {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	id_product := idProduct.ID

	_, err = h.productService.DeleteProduct(id_product)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	response := helper.APIResponse("success", "Product has been successfully deleted")
	c.JSON(http.StatusOK, response)
}
