package product

import "github.com/scys12/modul-go/internal/model"

func (pr *ProductRepo) CheckUsernameAndEmail(username, email string) (isExist bool, err error) {
	var count int
	const query = "SELECT COUNT(1) FROM product WHERE username=? OR email=?"
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
	const query = "INSERT INTO product(name, price, description, seller_id, is_bought) VALUES (?, ?, ?, ?, ?)"
	err = pr.dbContext.QueryRow(query, product.Name, product.Price, product.SellerID, product.IsBought).Err()
	return
}

func (pr *ProductRepo) GetCurrentUserProduct(id int, sellerID int) (product model.Product, err error) {
	const query = "SELECT * FROM product WHERE id=? and seller_id=?"
	err = pr.dbContext.QueryRow(query, id, sellerID).Scan(&product.ID, &product.Name, &product.Price,
		&product.Description, &product.IsBought, &product.SellerID, &product.UpdatedAt, &product.CreatedAt)
	return
}

func (pr *ProductRepo) GetProduct(id int) (product model.Product, err error) {
	const query = "SELECT * FROM product WHERE id=?"
	err = pr.dbContext.QueryRow(query, id).Scan(&product)
	return
}

func (pr *ProductRepo) UpdateProduct(product model.Product) (err error) {
	const query = "UPDATE product SET name=?, price=?, description=?, is_bought=? FROM  WHERE id=?"
	err = pr.dbContext.QueryRow(query, product.Name, product.Price, product.Description, product.IsBought, product.ID).Err()
	return
}

func (pr *ProductRepo) DeleteProduct(id int, sellerID int) (err error) {
	const query = "DELETE FROM product WHERE id=?, seller_id=?"
	err = pr.dbContext.QueryRow(query, id, sellerID).Err()
	return
}

func (pr *ProductRepo) GetSellerProducts(id int) (products []model.Product, err error) {
	const query = "SELECT * FROM product WHERE seller_id=?"
	rows, err := pr.dbContext.Query(query, id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		newProduct := model.Product{}
		err = rows.Scan(&newProduct.ID, &newProduct.Name, &newProduct.Price, &newProduct.SellerID, newProduct.Description, &newProduct.IsBought)
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
	const query = "UPDATE product SET is_bought=1 WHERE id=? and is_bought=0"
	err = pr.dbContext.QueryRow(query, id).Err()
	return
}

func (pr *ProductRepo) GetAvailableProducts() (products []model.Product, err error) {
	const query = "SELECT * FROM product WHERE is_bought=0"
	rows, err := pr.dbContext.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		newProduct := model.Product{}
		err = rows.Scan(&newProduct.ID, &newProduct.Name, &newProduct.Price, &newProduct.SellerID, newProduct.Description, &newProduct.IsBought)
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
