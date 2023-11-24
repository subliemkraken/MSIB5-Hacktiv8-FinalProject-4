package input

type CategoryCreateInput struct {
	Type              string `json:"type" valid:"required"`
	SoldProductAmount int    `json:"sold_product_amount"`
}

type CategoryUpdateInput struct {
	Type string `json:"type" valid:"required"`
}

type CategoryUpdateID struct {
	ID int `uri:"id" valid:"required"`
}

type CategoryDeleteID struct {
	ID int `uri:"id" valid:"required"`
}
