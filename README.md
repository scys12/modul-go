# Modul Golang

Pre Requirements:

-   Kemampuan Dasar Bahasa Pemrograman C.

-   Menginstall Visual Studio Code

# Modul Hari Pertama
## Install Golang (approx +- 20 min)

1.  Buka [https://golang.org/doc/install](https://www.google.com/url?q=https://golang.org/doc/install&sa=D&source=editors&ust=1617499705407000&usg=AOvVaw3yR2qwJ_3zpWD8rUUq0fiK)

![](img/1-1.png)

2. Klik Download Go For Windows(Bagi yang OS Windows)
![](img/1-2.png)
   - Sesuaikan lokasi file dengan yang digambar, Lalu plih **Next** dan **Install**.
![](img/1-3.png)
   - Tampilan Go telah selesai diinstall.
   - Cek apakah go telah terinstall dengan menggunakan command `$ go version`.


3.  Instalasi Linux
    -  sudo wget [https://golang.org/dl/go1.16.2.linux-amd64.tar.gz](https://www.google.com/url?q=https://golang.org/dl/go1.16.2.linux-amd64.tar.gz&sa=D&source=editors&ust=1617499705409000&usg=AOvVaw0U-IXHQZep3CWl_NcemuLF)  
    -  `rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz`  
    -  Buka file ~/.profile, lalu tambahkan baris berikut ke akhir file :  
![](img/1-4.png)
    -  Simpan lalu restart terminal anda.  
    -  Cek apakah go telah terinstall dengan menggunakan command `$ go version`.

4.  Instalasi Mac
![](img/1-5.png)
    - Pilih tab **Mac**  sesuai gambar.  
    - Download file dengan **Download Go For Mac**.  
    - Jangan lupa menambahkan path environment Go yaitu /usr/local/go/bin ke PATH environment variable  
    - Cek apakah go telah terinstall dengan menggunakan command `$ go version`.

## Integrasi Visual Studio Code dengan Go
![](img/1-6.png)
1.  Install plugin sesuai gambar diatas. Lalu restart VS Code kalian.
![](img/1-7.png)
2.  Jika disuruh menginstall plugin pada Go, maka pilih opsi Install All.

## Syntax Dasar dan Cara Menjalankan Go

![](img/1-8.png)
1.  Buat sebuah folder baru dan bukalah folder tersebut di **VSCode** kalian. Pada gambar diatas, nama folder adalah modul-go.
![](img/1-9.png)
2.  Untuk membuat suatu projek berbasis Go baru tanpa harus membuat mengaturnya pada $GOPATH, anda harus menginisiasi sebuah module dimana module ini berperan sebagai penanda bahwa di folder tersebut merupakan sebuah module untuk projek Go. Hal ini sangat berperan penting dalam pengembangan Go, karena jika kalian tidak membuat module pada folder baru ini, kalian wajib membuat folder pada lokasi file Go diinstall, dan itu membuat konfigurasi yang sangat susah.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Dalam hal ini kita menjalankan perintah `go mod init example.com/m`. Jika melihat lebih dalam lagi untuk penamaan nama module biasanya mengikuti {nama_domain}/{tim/user_pembuat}/{nama_project}. Pada contoh diatas dapat dilihat bahwa nama modulnya **github.com/scys12/modul-go** dimana modul tersebut berada di github.com dengan nama pembuat scys12 dan modul bernama modul-go. Nama modul ini berarti sebagai identifier, yang akan dibahas lebih lanjut pada topik selanjutnya.

3.  Ketika anda telah menginisiasi module, maka akan dibuat file baru bernama **go.mod**

![](img/1-10.png)

4.  Buatlah file **main.go**.  

Struktur file go biasanya wajib memiliki **package {package_name}** dimana nama package itu menandakan file tersebut berada di folder yang mana. Ketika ingin menggunakan library yang tersedia pada Go, penulisannya seperti berikut `Import “library_name”`, jika hanya satu library. Jika lebih dari satu library,
````GO
    Import (
    “Library_name1”
    “library_name2”
    )
````
Tapi tenang saja, karena ketika anda menggunakan plugin VSCode Go, maka plugin tersebut akan mengatur penulisan file Go anda.
-   Agar suatu file go berjalan wajib memiliki
**package main**
Dengan nama fungsi
````GO
func main (){
...
}
````

-   Jika suatu file memiliki func main(), maka Go akan mengeksekusi programnya melalui fungsi tersebut.![](img/1-11.png)

5.  Untuk mengeksekusi suatu projek Go, anda dapat melakukannya dengan perintah
``go build``  
```./{root_name_folder}```

Jika melihat dari gambar diatas, hasil dari perintah build akan menampilkan file baru yang namanya sesuai dengan base root folder, dimana dari gambar tersebut adalah modul-go. Jika anda menggunakan windows, dapat menjalankan program dengan perintah :

 ``{root_name_folder.exe}``

Lalu akan ditampilkan hasil dari program anda. Dari gambar diatas, program kita menampilkan

    Hello, 世界

## Syntax Dasar dan Pemahaman Go

###  Fungsi

![](img/1-12.png)

- Untuk membuat fungsi pada Go, anda dapat menggunakan syntax
````GO
func nama_fungsi(data_type param_name, data_type param_name2,...) return_type {
...
}
````

- Untuk fungsi yang tidak mengembalikan suatu nilai, maka return_type dapat dikosongi. Pada contoh diatas, membuat fungsi bernama add dengan parameter x dan y yang mengembalikan x+y dimana x+y merupakan integer.

   Untuk mengeksekusi agar hasil penjumlahan tersebut terlihat dapat menjalankan :

```go build```  
``./{root_folder_name}``

![](img/1-13.png)

Contoh lain:

![](img/1-14.png)

![](img/1-15.png)

###  Variable

![](img/1-16.png)

- Ketika ingin mendefinisikan variable dapat menggunakan syntax

````GO
var variable_name data_type
````

Namun variable ini belum didefinisikan nilainya. Atau kalian bisa dengan menggunakan
````GO
var variable_name = value
````
Dimana kita tidak perlu menuliskan type data, namun wajib memberi assignment terhadap variable tersebut.

Dapat dilihat, terdapat error yang muncul karena kita mendefinisikan suatu variable namun tidak menggunakannya. Go merupakan bahasa yang sensitif terhadap variablenya, jika ada variable yang didefinisikan tapi tidak digunakan, maka akan menampilkan error, sehingga cara terbaik yaitu dengan memastikan bahwa variable tersebut diassignment nilainya.

![](img/1-17.png)

- Terdapat suatu cara untuk mendefinisikan variable beserta valuenya yaitu dengan menggunakan tanda “**:=**”
````GO
var_name := value
````
### Flow Control

![](img/1-18.png)

Ketika ingin menggunakan **for** dalam Go, anda dapat mendefinisikannya seperti gambar diatas.

![](img/1-19.png)

Sedangkan untuk penggunaan **while**, Go tidak menyediakan syntax tersebut namun digantikan dengan **for** seperti gambar diatas.

![](img/1-20.png)

Terdapat **if/else** yang merupakan percabangan atau condition blocks. **If** akan bergantung pada hasil yang ingin dibandingkan.

### Array

![](img/1-21.png)

Apa itu **Array**? **Array** adalah sebuah struktur data yang terdiri atas banyak variabel dengan tipe data sama, dimana masing-masing elemen variabel mempunyai nilai indeks. Kalian dapat membuat fixed-length array dengan menetapkan berapa ukuran dari array, dengan syntax
````GO
[size_n]data_type{val1, val2, val3,..., valn}
````
Kalian juga dapat membuat array dynamic dimana kalian tidak perlu menetapkan ukuran dari array tersebut. Salah satu cara mendeklarasikan array dynamic :
````GO
Var variable_name [ ]data_type
````
Dan kalian dapat menambahkan value dari array tersebut dengan fungsi **append()**.

### Struct
Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti **map**, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda. Struct mirip dengan **Class** di **OOP**.

![](img/1-22.png)

Dari kodingan diatas, kita mempunyai **Struct** `Book`  yang mempunyai field `ISBN`, `Genre`, `Name` dengan tipe String. Lalu kita mempunyai **Struct** `Library` dengan field `Name` (string), `Size`(int), Lalu kita mempunyai array of struct dari **Struct** `Book`.

![](img/1-23.png)

Kalian dapat mendeklarasikan struct seperti kodingan diatas.

`Book := Book{}` dapat dikatakan bahwa kita membuat suatu struk baru. Jika dikonversikan ke OOP Class, maka :

    Book book = Book()

Dimana kita tidak menaruh parameter pada konstruktornya

![](img/1-24.png)

Pada deklarasi struk diatas, kita mendeklarasi sekaligus mengisi value dari field field pada struct `Book`. Jika terdapat suatu field yang tidak diisi, maka default valuenya bisa berupa **nil**, **0**. **Nil** pada **Go** berarti  suatu variable belum memiliki value sama sekali. **Nil** mirip dengan **Null** pada bahasa pemrograman lainnya. Jika kodingan diatas dikonversikan ke OOP Class, maka :

    Book bookV2 = Book(“Horror”, “124”, “Foo”)

Dimana pada class Booknya menjadi :
````Java
class Book {
	Book(string isbn, string genre, string name){
		This.isbn = isbn
		This.genre = genre
		This.name = name
	}
}
````
![](img/1-25.png)

Untuk mendeklarasikan struct `Library`, kalian dapat membuat seperti kode diatas, atau

![](img/1-26.png)

Perbedaannya adalah pada gambar pertama, struct buku sudah kita assign ke variable book dan bookV2, sedangkan gambar kedua, kita mengisi `Book` tanpa mengassign ke variable.

### Defer

Defer merupakan syntax untuk mengeksekusi statement pada saat fungsi melakukan pengembalian.

![](img/1-27.png)

Kita mempunyai fungsi bernama DeferFunc() dengan return value berupa **boolean**.

Kegunaan defer adalah untuk memastikan suatu statement benar benar dijalankan. Dapat dilihat bahwa disana kita mengecek genre bukunya. Bagaimana kita bisa memastikan bahwa sebelum fungsi melakukan return, dia akan mengeksekusi suatu statement? Kita dapat memastikan dengan menggunakan defer.

![](img/1-28.png)

Output dari fungsi diatas.

### Method
![](img/1-29.png)

Apa itu method ? Method sama seperti function, namun yang membedakan method mempunyai receiver argument, dimana receiver argumen itu menandakan method tersebut dimiliki oleh struktur data yang mana.

Dari gambar diatas kita mempunyai struct `User`. Lalu kita mempunyai 3 Method dasri struct User yaitu **GetUserNameAndId()**, **generatePassword()**, **GetPassword()**.

Kenapa 3 fungsi ini dibilang method, karena mereka mempunyai receiver, yaitu (u user).

````GO
Func (receiver_name receiver_type) function_name(param1,...,paramn
````
Statement diatas adalah cara membuat method.

### Public/Private/Protected Identifier
Mungkin kalian bertanya-tanya, gimana caranya membuat public/private/protected seperti yang ada di bahasa pemrograman OOP lainnya. Di Go,tidak ada terminologi yang menerangkan bagaimana membuat public/private/protected, melainkan dengan menggunakan huruf besar atau huruf kecil. Agar suatu variable/struct/fungsi/method dianggap sebagai public identifier, harus menggunakan huruf besar pada awalan nama. Sehingga kita bisa memanggil dari **package** yang berbeda. Sedangkan jika menggunakan huruf kecil pada awalan nama, maka akan dianggap protected dan private. Lebih lanjut suatu variable/struct/fungsi/method dianggap private ketika dia tidak dapat diakses dari **package diluar package dari file tersebut (didefinisikan dibaris paling atas).** Sedangkan protected dianggap ketika kita bisa mengakses dari package yang sama dengan package file. Jadi walaupun terdapat 2 file berbeda namun masih satu package, maka kita bisa mengakses variable/struct/fungsi/method karena dianggap protected.

![](img/1-30.png)

Struct `User` berada di package `user`, dan kita ingin memanggil melalui package **main**. Dapat dilihat bahwa, fungsi yang muncul hanya fungsi berawalan huruf besar, sedangkan fungsi yang berawalan huruf kecil tidak dapat diakses dari luar package.

### Pointer

Pointer menyimpan alamat-alamat pada Go. Pointer sangat penting ketika menggunakan Go.

![](img/1-31.png)

![](img/1-32.png)

Kita dapat mendeklarasikan variable yang menggunakan pointer seperti kodingan diatas. Kita dapat menggunakan new(), atau menggunakan tanda “&” pada waktu deklarasi.

![](img/1-33.png)

![](img/1-34.png)

Pointer juga dapat digunakan sebagai parameter, dimana variable yang kita passing juga harus dalam bentuk pointer. Jika dilihat bahwa `user` sudah dalam bentuk pointer (Gambar Sebelumnya)  sehingga jika kita passing ke fungsi **ChangeUserData()** maka dapat berjalan. Dapat dilihat bahwa fungsi **ChangeUserData**, kita hanya perlu mengganti isi dari struct kita tanpa melakukan return.

Jika kita awalnya mendeklarasikan variable tanpa menggunakan pointer yaitu variable `userV3`, ketika kita ingin passing ke fungsi **ChangeUserData()** maka dapat menggunakan tanda “&” diparameter fungsi tersebut.

![](img/1-35.png)

![](img/1-36.png)

Ketika kita ingin mempassing suatu variable berpointer ke dalam fungsi yang tidak menggunakan pointer, maka kita dapat menggunakan tanda “*” di variable yang ingin dipassing
### Pointer Receiver.
![](img/1-37.png)

Pada kode diatas kita membuat receiver menggunakan pointer, yaitu `(u *User)`. Ada 2 alasan kenapa kita harus menggunakan pointer receiver.

Yang pertama adalah agar method dapat mengubah value dari field2 di receivernya

Yang kedua adalah menghindari menyalin nilai pada setiap panggilan metode. Ini bisa lebih efisien jika receiver adalah sebuah struct besar
>**Rule Of Thumb:** Ketika kalian tidak yakin terhadap apa yang akan receiver kalian lakukan, gunakan pointer.