package main

import (
	"errors"
	"fmt"
	"strconv"
)

//9 Interface
type StudentBehavior interface {
	PayFee()
	GoToSchool() bool
	TalkToFriends(friend_name string) (string, string)
}

//9 Interface
type Student struct {
	ID   int
	Name string
}

//Mengimplementasi Interface

//9 Interface
func (sb *Student) PayFee() {
	panic("Implement me")
}

//9 Interface
func (sb *Student) GoToSchool() bool {
	return true
}

//9 Interface
func (sb *Student) TalkToFriends(friend_name string) (string, string) {
	return fmt.Sprintf("From %v", sb.Name), fmt.Sprintf("To %v", friend_name)
}

//9 Interface
func InitStudent() {
	var studentBehave StudentBehavior = &Student{
		ID:   1,
		Name: "Sam",
	}
	fmt.Println(studentBehave.GoToSchool())
	fmt.Println(studentBehave.TalkToFriends("Ivan"))

}

//9 Empty Interface
func InitEmptyInterface() {
	var value interface{}
	value = 42
	fmt.Println(value)
	value = "abcd"
	fmt.Println(value)
	value = &Student{
		ID:   2,
		Name: "Simon",
	}
	fmt.Println(value)
}

func InitMaps() {
	var mapStrInt map[string]int
	mapStrInt = make(map[string]int)
	mapStrInt["1"] = 1
	mapStrInt["2"] = 2
	mapStrInt["3"] = 3

	var mapStrStudent map[string]*Student
	mapStrStudent = make(map[string]*Student)
	mapStrStudent["teladan"] = &Student{ID: 10, Name: "Samuel"}
	mapStrStudent["terburuk"] = &Student{ID: 11, Name: "Wahed"}
}

func InitErrors() error {
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return err
	}
	fmt.Println("Converted integer:", i)
	if i < 10 {
		err := errors.New("Angkanya sedikit")
		return err
	}
	return nil
}
