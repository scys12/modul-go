package model

type Product struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	SellerID    int    `json:"seller_id"`
	IsBought    bool   `json:"is_bought"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
