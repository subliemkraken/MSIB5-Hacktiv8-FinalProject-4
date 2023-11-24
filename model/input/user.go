package input

type UserRegisterInput struct {
	FullName string `json:"full_name" valid:"required"`
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required,minstringlength(6)"`
	Role     string `json:"role"`
	Balance  int    `json:"balance"`
}

type UserLoginInput struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
}

type UserUpdateInput struct {
	FullName string `json:"full_name" `
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" `
	Role     string `json:"role"`
	Balance  int    `json:"balance"`
}

type UserTopupInput struct {
	Balance int `json:"balance" valid:"required"`
}

type UserUpdateID struct {
	ID int `uri:"id" valid:"required"`
}

type UserDeleteID struct {
	ID int `uri:"id" valid:"required"`
}
