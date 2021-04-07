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

type ProductRequest struct {
	Name        string `json:"name" validate:"required,max=50"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required,max=500"`
}
