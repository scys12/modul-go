package user

import (
	"encoding/base64"
	"fmt"
)

type User struct {
	Id       int
	Name     string
	Password string
}

//public method
func (u User) GetUserNameAndId() string {
	return fmt.Sprintf("%v %v", u.Id, u.Name)
}

//private method / protected method
func (u User) generatePassword() string {
	return base64.StdEncoding.EncodeToString([]byte(u.Password))
}

//public method
func (u User) GetPassword() string {
	return u.generatePassword()
}

func (u *User) ChangeData() {
	u.Id = 10
	u.Name = "Sam"
	u.Password = "Password123"
}

func InitUser() {
	//8 Pointer
	user := &User{
		Id:       12,
		Name:     "Sam",
		Password: "123456",
	}

	fmt.Printf("%T", user)

	//8 Pointer
	userV1 := &User{}
	fmt.Printf("%T", userV1)

	//8 Pointer
	userV2 := new(User)
	fmt.Printf("%T", userV2)

	ChangeUserData(user)

	fmt.Println(user.Id)
	fmt.Println(user.Name)
	fmt.Println(user.Password)

	userV3 := User{
		Id:       1,
		Name:     "Sam",
		Password: "123456789",
	}
	ChangeUserData(&userV3)
	fmt.Println(user.Id)
	fmt.Println(user.Name)
	fmt.Println(user.Password)

	userV3 = ChangeUserDataWithReturn(*user)

	fmt.Println(user.Id)
	fmt.Println(user.Name)
	fmt.Println(user.Password)
}

//8 Pointer in parameter
func ChangeUserData(usr *User) {
	usr.Id = 45
	usr.Name = "Samuel"
	usr.Password = "Password12345"
}

func ChangeUserDataWithReturn(usr User) User {
	usr.Id = 4
	usr.Name = "Samuel1"
	usr.Password = "Password123"
	return usr
}
