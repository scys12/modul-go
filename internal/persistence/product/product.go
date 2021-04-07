package product

import (
	"github.com/scys12/modul-go/internal/model"
)

func (pr *ProductRepo) CheckUsernameAndEmail(username, email string) (isExist bool, err error) {
	var count int
	const query = "SELECT COUNT(1) FROM products WHERE username=? OR email=?"
	isExist = false
	err = pr.dbContext.QueryRow(query, username, email).Scan(&count)
	if err != nil {
		return
	}
	if count > 0 {
		isExist = true
	}
	return
}

func (pr *ProductRepo) InsertProduct(product model.Product) (err error) {
	const query = "INSERT INTO products(name, price, description, seller_id) VALUES (?, ?, ?, ?)"
	err = pr.dbContext.QueryRow(query, product.Name, product.Price, product.Description, product.SellerID).Err()
	return
}

func (pr *ProductRepo) GetCurrentUserProduct(id int, sellerID int) (product model.Product, err error) {
	const query = "SELECT * FROM products WHERE id=? and seller_id=?"
	err = pr.dbContext.QueryRow(query, id, sellerID).Scan(&product.ID, &product.Name, &product.Price,
		&product.Description, &product.IsBought, &product.SellerID, &product.UpdatedAt, &product.CreatedAt)
	return
}

func (pr *ProductRepo) GetProduct(id int) (product model.Product, err error) {
	const query = "SELECT * FROM products WHERE id=?"
	err = pr.dbContext.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.IsBought, &product.SellerID, &product.CreatedAt, &product.UpdatedAt)
	return
}

func (pr *ProductRepo) UpdateProduct(product model.Product) (err error) {
	const query = "UPDATE products SET name=?, price=?, description=? WHERE id=?"
	err = pr.dbContext.QueryRow(query, product.Name, product.Price, product.Description, product.ID).Err()
	return
}

func (pr *ProductRepo) DeleteProduct(id int, sellerID int) (err error) {
	const query = "DELETE FROM products WHERE id=? and seller_id=?"
	err = pr.dbContext.QueryRow(query, id, sellerID).Err()
	return
}

func (pr *ProductRepo) GetSellerProducts(id int) (products []model.Product, err error) {
	products = []model.Product{}
	const query = "SELECT * FROM products WHERE seller_id=?"
	rows, err := pr.dbContext.Query(query, id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		newProduct := model.Product{}
		err = rows.Scan(&newProduct.ID, &newProduct.Name, &newProduct.Price, &newProduct.Description, &newProduct.IsBought, &newProduct.SellerID, &newProduct.CreatedAt, &newProduct.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, newProduct)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return
}

func (pr *ProductRepo) BuyProduct(id int) (err error) {
	const query = "UPDATE products SET is_bought=1 WHERE id=? and is_bought=0"
	err = pr.dbContext.QueryRow(query, id).Err()
	return
}

func (pr *ProductRepo) GetAvailableProducts() (products []model.Product, err error) {
	products = []model.Product{}
	const query = "SELECT * FROM products WHERE is_bought=0"
	rows, err := pr.dbContext.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		newProduct := model.Product{}
		err = rows.Scan(&newProduct.ID, &newProduct.Name, &newProduct.Price, &newProduct.Description, &newProduct.IsBought, &newProduct.SellerID, &newProduct.CreatedAt, &newProduct.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, newProduct)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return
}
