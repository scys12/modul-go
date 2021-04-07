package user

import (
	"github.com/scys12/modul-go/internal/model"
)

func (ur *UserRepo) CheckUsernameAndEmail(username, email string) (isExist bool, err error) {
	var count int
	const query = "SELECT COUNT(1) FROM users WHERE username=? OR email=?"
	isExist = false
	err = ur.dbContext.QueryRow(query, username, email).Scan(&count)
	if err != nil {
		return
	}
	if count > 0 {
		isExist = true
	}
	return
}

func (ur *UserRepo) InsertUser(user model.User) (err error) {
	const query = "INSERT INTO users(username, password, email) VALUES (?, ?, ?)"
	err = ur.dbContext.QueryRow(query, user.Username, user.Password, user.Email).Err()
	return
}

func (ur *UserRepo) GetUser(username string) (user model.User, err error) {
	const query = "SELECT * FROM users WHERE username=?"
	err = ur.dbContext.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return
}

func (ur *UserRepo) GetUserByID(id int) (user model.User, err error) {
	const query = "SELECT * FROM users WHERE id=?"
	err = ur.dbContext.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return
}
