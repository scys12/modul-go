package model

type Seller struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
