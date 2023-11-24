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

type transactionHistoryController struct {
	transactionHistoryService service.TransactionHistoryService
	productService            service.ProductService
}

func NewTransactionHistoryController(transactionHistoryService service.TransactionHistoryService, productService service.ProductService) *transactionHistoryController {
	return &transactionHistoryController{transactionHistoryService, productService}
}

func (h *transactionHistoryController) CreateTransactionHistory(c *gin.Context) {
	var input input.TransactionHistoryCreateInput

	currentUser := c.MustGet("currentUser").(int)

	err := c.ShouldBindJSON(&input)

	transactionHistory, err := govalidator.ValidateStruct(input)

	product, err := h.productService.FindProductByID(input.ProductID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	priceProduct := product.Price

	if !transactionHistory {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	result, err := h.transactionHistoryService.CreateTransactionHistory(input, currentUser, priceProduct)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	transactionHistoryResponse := response.TransactionHistoryCreateResponse{
		TotalPrice:   result.TotalPrice,
		Quantity:     result.Quantity,
		ProductTitle: product.Title,
	}

	message := "You have succesfully purchased the product"

	response := helper.TransAPIResponse("created", message, transactionHistoryResponse)
	c.JSON(http.StatusCreated, response)
}

func (h *transactionHistoryController) GetMyTransactionHistory(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	transactionHistory, err := h.transactionHistoryService.GetMyTransactionHistory(currentUser)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	response := helper.APIResponse("success", transactionHistory)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHistoryController) GetAllTransactionHistory(c *gin.Context) {
	transactionHistory, err := h.transactionHistoryService.GetAllTransactionHistory()

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	response := helper.APIResponse("success", transactionHistory)
	c.JSON(http.StatusOK, response)
}
