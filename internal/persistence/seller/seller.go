package seller

import "github.com/scys12/modul-go/internal/model"

func (sr *SellerRepo) CheckUsernameAndEmail(username, email string) (isExist bool, err error) {
	var count int
	const query = "SELECT COUNT(1) FROM seller WHERE username=? OR email=?"
	isExist = false
	err = sr.dbContext.QueryRow(query, username, email).Scan(&count)
	if err != nil {
		return
	}
	if count > 0 {
		isExist = true
	}
	return
}

func (sr *SellerRepo) InsertSeller(seller model.Seller) (err error) {
	const query = "INSERT INTO seller(username, password, email) VALUES (?, ?, ?)"
	err = sr.dbContext.QueryRow(query, seller.Username, seller.Password, seller.Email).Err()
	return
}

func (sr *SellerRepo) GetSeller(username string) (seller model.Seller, err error) {
	const query = "SELECT * FROM seller WHERE username=?"
	err = sr.dbContext.QueryRow(query, username).Scan(&seller.ID, &seller.Username, &seller.Password, &seller.Email, &seller.CreatedAt, &seller.UpdatedAt)
	return
}

func (sr *SellerRepo) GetCustomerByID(id int) (seller model.Seller, err error) {
	const query = "SELECT * FROM seller WHERE id=?"
	err = sr.dbContext.QueryRow(query, id).Scan(&seller.ID, &seller.Username, &seller.Password, &seller.Email, &seller.CreatedAt, &seller.UpdatedAt)
	return
}
