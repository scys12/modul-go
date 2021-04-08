package main

import (
	"fmt"

	"github.com/scys12/modul-go/user"
)

func main() {
	First()
	Second()
	DeferFunc()

	usr := user.User{}
	usr.Id = 500
	usr.Name = "Wahed"
	usr.Password = "Wahed123"
	text := usr.GetUserNameAndId() //public function
	fmt.Println(text)
	pass := usr.GetPassword()
	fmt.Println(pass)

	InitStudent()
	InitEmptyInterface()
}
