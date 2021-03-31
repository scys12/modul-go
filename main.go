package main

import "fmt"

// Fungsi add mempunyai paramater x (integer) & y (integer) dengan
// return type int
func add(x int, y int) int {
	return x + y
}

// Fungsi swap mempunyai paramater x (string) & y(string) dengan
//return type berjumlah 2, yaitu string dan string
func swap(x, y string) (string, string) {
	return y, x
}

// Fungsi split mempunyai parameter sum (integer) dengan
//return type yang sudah didefenisi variabelnya sebelumnya
//sehingga waktu menuliskan return, kita tidak perlu menuliskan
//secara lengkap apa yg direturn, misal : return x,y tidak diperlukan
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var a, b, c bool //3. Variable

func main() {
	fmt.Println("Hello, 世界") // 1. Intro

	fmt.Println(add(42, 13))      //2. Fungsi
	fmt.Println(swap("42", "13")) //2. Fungsi
	fmt.Println(split(17))        //2. Fungsi

	var i int                       //3. Variable
	var k, l int = 1, 2             //3. Variable
	var y, x, z = true, false, "no" //3. Variable

	o := 3                                  //3. variable
	log, python, java := true, false, "no!" //3.variable

	sum := 0                  //4. Flow Control
	for i := 0; i < 10; i++ { // 4. Flow Control
		sum += i
	}
	fmt.Println(sum) // 4. Flow Control

	sum = 1          //4. Flow Control
	for sum < 1000 { //4. Flow Control
		sum += sum
	}
	fmt.Println(sum) //4 Flow Control
}
