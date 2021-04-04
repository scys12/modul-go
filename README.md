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