package input

type ProductCreateInput struct {
	ID         int    `json:"id"`
	Title      string `json:"title" valid:"required"`
	Price      int    `json:"price" valid:"required"`
	Stock      int    `json:"stock" valid:"required"`
	CategoryID int    `json:"category_id" valid:"required"`
}

type ProductUpdateInput struct {
	Title      string `json:"title" valid:"required"`
	Price      int    `json:"price" valid:"required"`
	Stock      int    `json:"stock" valid:"required"`
	CategoryID int    `json:"category_id" valid:"required"`
}

type ProductUpdateID struct {
	ID int `uri:"id" valid:"required"`
}

type ProductDeleteInput struct {
	ID int `uri:"id" valid:"required"`
}
