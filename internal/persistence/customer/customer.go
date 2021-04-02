package customer

import (
	"github.com/scys12/modul-go/internal/model"
)

func (cr *CustomerRepo) CheckUsernameAndEmail(username, email string) (isExist bool, err error) {
	var count int
	const query = "SELECT COUNT(1) FROM customer WHERE username=? OR email=?"
	isExist = false
	err = cr.dbContext.QueryRow(query, username, email).Scan(&count)
	if err != nil {
		return
	}
	if count > 0 {
		isExist = true
	}
	return
}

func (cr *CustomerRepo) InsertCustomer(customer model.Customer) (err error) {
	const query = "INSERT INTO customer(username, password, email) VALUES (?, ?, ?)"
	err = cr.dbContext.QueryRow(query, customer.Username, customer.Password, customer.Email).Err()
	return
}

func (cr *CustomerRepo) GetCustomer(username string) (customer model.Customer, err error) {
	const query = "SELECT * FROM customer WHERE username=?"
	err = cr.dbContext.QueryRow(query, username).Scan(&customer.ID, &customer.Email, &customer.Password, &customer.Username, &customer.UpdatedAt, &customer.CreatedAt)
	return
}

func (cr *CustomerRepo) GetCustomerByID(id int) (customer model.Customer, err error) {
	const query = "SELECT * FROM customer WHERE id=?"
	err = cr.dbContext.QueryRow(query, id).Scan(&customer.ID, &customer.Email, &customer.Password, &customer.Username, &customer.UpdatedAt, &customer.CreatedAt)
	return
}
