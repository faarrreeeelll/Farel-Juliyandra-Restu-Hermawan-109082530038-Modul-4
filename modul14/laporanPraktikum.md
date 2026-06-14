# <h1 align="center">Laporan Praktikum Modul 14 - ... </h1>
<p align="center">[Farel Juliyandra Restu Hermawan] - [109082530038]</p>

## Unguided 


### 1. [Soal 1 14.3]
#### soal1-14.3.go

```go
package main

import "fmt"

const NMax int = 1000000

type ArrayRumah [NMax]int

func SelectionSort(arr *ArrayRumah, n int) {
	var i, j, minIdx, temp int

	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		temp = arr[minIdx]
		arr[minIdx] = arr[i]
		arr[i] = temp
	}
}

func main() {
	var n, m, i, j int
	var rumah ArrayRumah

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&m)

		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		SelectionSort(&rumah, m)

		for j = 0; j < m; j++ {
			fmt.Print(rumah[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul14/output/output-soal1-14.3%20.png)
[Program ini membaca jumlah daerah ($n$) dan jumlah rumah ($m$) beserta rangkaian nomor rumah para kerabat di setiap daerah tersebut. Program kemudian mengurutkan nomor-nomor rumah di setiap daerah secara terurut membesar menggunakan algoritma selection sort. Setelah proses pengurutan selesai, program akan langsung menampilkan rangkaian nomor rumah yang sudah rapi tersebut dalam satu baris per daerahnya.]


### 2. [Soal 1 14.3]
#### soal2-14.3.go

```go
package main

import "fmt"

const NMax int = 1000000

type ArrayRumah [NMax]int

func SelectionSort(arr *ArrayRumah, n int) {
	var i, j, minIdx, temp int

	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		temp = arr[minIdx]
		arr[minIdx] = arr[i]
		arr[i] = temp
	}
}

func main() {
	var n, m, i, j int
	var rumah ArrayRumah
	var pertama bool

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&m)

		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		SelectionSort(&rumah, m)

		pertama = true

		for j = 0; j < m; j++ {
			if rumah[j]%2 != 0 {
				if !pertama {
					fmt.Print(" ")
				}
				fmt.Print(rumah[j])
				pertama = false
			}
		}

		for j = m - 1; j >= 0; j-- {
			if rumah[j]%2 == 0 {
				if !pertama {
					fmt.Print(" ")
				}
				fmt.Print(rumah[j])
				pertama = false
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul14/output/output-soal2-14.3%20.png)
[Program ini membaca data daerah dan nomor rumah kerabat dengan format yang sama seperti program pertama, lalu mengurutkan seluruh data tersebut secara membesar menggunakan algoritma selection sort. Pada saat pencetakan keluaran, program memisahkan data dengan cara menampilkan semua nomor rumah ganjil terlebih dahulu (dari terkecil ke terbesar), kemudian langsung diikuti dengan menampilkan semua nomor rumah genap (dari terbesar ke terkecil) dalam satu baris yang sama.]


### 3. [Soal 1 14.3]
#### soal1-14.3.go

```go
package main

import "fmt"

const NMax int = 1000000

type ArrayData [NMax]int

func InsertionSort(arr *ArrayData, n int) {
	var i, j, key int
	for i = 1; i < n; i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var data ArrayData
	var n, val int

	n = 0
	for {
		fmt.Scan(&val)

		if val == -5313 || val < 0 {
			break
		}

		if val == 0 {
			if n > 0 {
				InsertionSort(&data, n)
				if n%2 != 0 {
					fmt.Println(data[n/2])
				} else {
					fmt.Println((data[n/2-1] + data[n/2]) / 2)
				}
			}
		} else {
			data[n] = val
			n++
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul14/output/output-soal3-14.3%20.png)
[Program ini membaca sekumpulan bilangan bulat secara terus-menerus hingga mendeteksi angka penanda negatif (-5313) untuk menghentikan program. Setiap kali program membaca angka 0, data yang telah tersimpan hingga saat itu akan diurutkan menggunakan algoritma insertion sort, lalu program akan mencari dan mencetak nilai mediannya (jika jumlah data ganjil diambil nilai tengahnya, jika genap diambil rata-rata dua nilai tengahnya yang dibulatkan ke bawah). Bilangan selain 0 dan angka penanda akan terus disimpan ke dalam array untuk perhitungan selanjutnya.]

### 1. [Soal 1 14.6]
#### soal1-14.6.go

```go
package main

import "fmt"

const NMAX int = 100

func insertionSort(data *[NMAX]int, n int) {
	var i, j, key int

	for i = 1; i < n; i++ {
		key = data[i]
		j = i - 1

		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = key
	}
}

func cekJarak(data [NMAX]int, n int) {
	var i, jarak int
	var tetap bool = true

	if n < 2 {
		fmt.Println("Data berjarak 0")
		return
	}

	jarak = data[1] - data[0]

	for i = 2; i < n; i++ {
		if data[i]-data[i-1] != jarak {
			tetap = false
		}
	}

	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

func main() {
	var data [NMAX]int
	var x, n int

	fmt.Scan(&x)

	for x >= 0 && n < NMAX {
		data[n] = x
		n++
		fmt.Scan(&x)
	}

	insertionSort(&data, n)

	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	cekJarak(data, n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul14/output/output-soal1-14.6.png)
[Program ini membaca sekumpulan bilangan bulat positif dari input (hingga nilai negatif), kemudian mengurutkannya menggunakan algoritma insertion sort. Setelah diurutkan, program mencetak data yang sudah terurut, lalu melakukan pengecekan jarak (selisih) antara elemen-elemen yang berurutan untuk menentukan apakah data memiliki jarak yang konsisten atau tidak. Jika semua selisih antara elemen berurutan sama, maka program menampilkan nilai jaraknya; jika tidak konsisten, maka program menampilkan pesan "Data berjarak tidak tetap".]

### 2. [Soal 2 14.6]
#### soal2-14.6.go

```go
package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var i, idxMax int

	idxMax = 0

	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}

	fmt.Println(pustaka[idxMax].judul, pustaka[idxMax].penulis, pustaka[idxMax].penerbit, pustaka[idxMax].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var i, j int
	var temp Buku

	for i = 1; i < n; i++ {
		temp = pustaka[i]
		j = i - 1

		for j >= 0 && pustaka[j].rating < temp.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int

	batas = 5
	if n < 5 {
		batas = n
	}

	for i = 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	var kiri, kanan, tengah int
	var ditemukan bool

	kiri = 0
	kanan = n - 1
	ditemukan = false

	for kiri <= kanan && !ditemukan {
		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == r {
			ditemukan = true
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ditemukan {
		fmt.Println(
			pustaka[tengah].judul,
			pustaka[tengah].penulis,
			pustaka[tengah].penerbit,
			pustaka[tengah].tahun,
			pustaka[tengah].eksemplar,
			pustaka[tengah].rating,
		)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka, ratingCari int

	fmt.Scan(&nPustaka)

	DaftarkanBuku(&pustaka, nPustaka)

	fmt.Scan(&ratingCari)

	CetakTerfavorit(pustaka, nPustaka)

	UrutBuku(&pustaka, nPustaka)

	Cetak5Terbaru(pustaka, nPustaka)

	CariBuku(pustaka, nPustaka, ratingCari)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul14/output/output-soal2-14.6.png)
[Program ini merupakan sistem manajemen perpustakaan yang mengelola koleksi buku. Program membaca data buku berupa id, judul, penulis, penerbit, jumlah eksemplar, tahun, dan rating, kemudian melakukan beberapa operasi: mencari dan menampilkan buku dengan rating tertinggi, mengurutkan seluruh buku berdasarkan rating secara descending menggunakan insertion sort, menampilkan 5 buku teratas setelah pengurutan, dan melakukan pencarian binary search untuk menemukan buku dengan rating tertentu yang dimasukkan pengguna.]
