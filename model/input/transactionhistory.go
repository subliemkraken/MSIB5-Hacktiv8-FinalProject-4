package input

type TransactionHistoryCreateInput struct {
	ProductID int `json:"product_id" valid:"required"`
	Quantity  int `json:"quantity" valid:"required"`
}
