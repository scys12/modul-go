package payload

type LoginRequest struct {
	Username string `json:"username" validate:"required,max=25,min=4"`
	Password string `json:"password" validate:"required,max=25,min=4"`
}

type RegisterRequest struct {
	Email    string `json:"email,omitempty" validate:"required,email,max=25"`
	Username string `json:"username,omitempty" validate:"required,max=25,min=4"`
	Password string `json:"password,omitempty" validate:"required,max=25,min=4"`
}

type SellerRegisterRequest struct {
	Email    string `json:"email,omitempty" validate:"required,email,max=25"`
	Username string `json:"username,omitempty" validate:"required,max=25,min=4"`
	Password string `json:"password,omitempty" validate:"required,max=25,min=4"`
	Name     string `json:"name,omitempty" validate:"required,max=50,min=10"`
}

type InsertProductRequest struct {
	Name        string `json:"name" validate:"required,max=50"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required,max=500"`
}

type UpdateProductRequest struct {
	Name        string `json:"name,omitempty" validate:"max=50"`
	Price       int    `json:"price,omitempty" validate:"numeric"`
	Description string `json:"description,omitempty" validate:"max=500"`
}
