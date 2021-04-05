package main

import "fmt"

//6 Struct
type Book struct {
	ISBN  string
	Genre string
	Name  string
}

//6 Struct
type Library struct {
	Name string
	Size int
	Book []Book
}

func Second() {

	//6 Deklarasi Struct
	book := Book{}
	book.Genre = "Horror"
	book.ISBN = "123"
	book.Name = "Foo"

	//6 Deklarasi Struct
	bookV2 := Book{
		ISBN:  "Horror",
		Genre: "124",
		Name:  "Foo",
	}
	bookV2.Genre = "Action"

	fmt.Println(bookV2.Genre)
	fmt.Println(bookV2.ISBN)
	fmt.Println(bookV2.Name)

	//6 Deklarasi Struct
	library := Library{}
	library.Name = "Foo"
	library.Size = 100
	library.Book = []Book{book, bookV2}

	for _, x := range library.Book {
		fmt.Println(x.Name)
		fmt.Println(x.Genre)
		fmt.Println(x.ISBN)
	}

	//6 Deklarasi Struct
	_ = Library{
		Name: "abc",
		Size: 100,
		Book: []Book{
			book,
			bookV2,
		},
	}

	_ = Library{
		Name: "Foo",
		Size: 100,
		Book: []Book{
			{
				ISBN: "abc",
			},
			{
				Genre: "Adventure",
			},
		},
	}
}

func DeferFunc() bool {
	fmt.Println("Defer Func")
	defer fmt.Println("Defer berjalan tepat sebelum fungsi selesai") // 7 Defer

	bookV2 := Book{
		ISBN:  "Horror",
		Genre: "124",
		Name:  "Foo",
	}

	if bookV2.Genre == "Adventure" {
		return false
	}

	if bookV2.Genre == "Action" {
		return true
	}

	fmt.Println(bookV2)
	fmt.Println("Fungsi selesai")
	return true
}
