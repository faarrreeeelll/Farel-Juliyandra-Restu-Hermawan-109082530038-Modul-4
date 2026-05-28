# <h1 align="center">Laporan Praktikum Modul 10 </h1>
<p align="center">[Farel Juliyandra Restu Hermawan] - [109082530038]</p>

## Unguided 

### 1. [Soal Latihan Modul 10]
#### soal1.go

```go
package main

import "fmt"

const N = 1000

func hitungMinMax(berat [N]float64, n int, min *float64, max *float64) {
	*min = berat[0]
	*max = berat[0]
	for i := 1; i < n; i++ {
		if berat[i] < *min {
			*min = berat[i]
		}

		if berat[i] > *max {
			*max = berat[i]
		}
	}
}

func main() {
	var berat [N]float64
	var max, min float64
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &min, &max)
	fmt.Printf("%.2f %.2f\n", min, max)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul10/output/output-soal1.png)
[Program di atas dibuat menggunakan bahasa Go untuk mencari nilai berat paling kecil dan paling besar dari sejumlah data yang dimasukkan pengguna. Pertama, program meminta jumlah data (n), lalu pengguna memasukkan data berat satu per satu ke dalam array berat. Setelah itu, fungsi hitungMinMax dipanggil untuk memeriksa seluruh isi array dan membandingkan setiap nilai agar diperoleh nilai minimum dan maksimum. Nilai terkecil disimpan pada variabel min, sedangkan nilai terbesar disimpan pada variabel max. Di akhir program, kedua nilai tersebut ditampilkan dengan format dua angka di belakang koma.]

### 2. [Soal Latihan Modul 10]
#### soal2.go

```go
package main

import "fmt"

const N = 1000

func main() {
	var berat [N]float64
	var x, y int
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	w := (x + y - 1) / y
	var tot [N]float64

	for i := 0; i < x; i++ {
		idx := i / y
		tot[idx] += berat[i]
	}

	for i := 0; i < w; i++ {
		fmt.Printf("%.2f", tot[i])
		if i < w-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	for i := 0; i < w; i++ {
		nIkan := y
		if i == w-1 && x%y != 0 {
			nIkan = x % y
		}
		rata := tot[i] / float64(nIkan)
		fmt.Printf("%.2f", rata)
		if i < w-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul10/output/output-soal2.png)
[Program di atas digunakan untuk mengelompokkan data berat ke dalam beberapa kelompok lalu menghitung total dan rata-rata setiap kelompok. Pertama, pengguna memasukkan jumlah data (x) dan jumlah anggota tiap kelompok (y), kemudian data berat disimpan ke dalam array berat. Program selanjutnya membagi data menjadi beberapa kelompok menggunakan perhitungan indeks, lalu menjumlahkan setiap kelompok ke array tot. Setelah semua data diproses, program menampilkan total berat tiap kelompok dengan format dua angka di belakang koma. Setelah itu, program juga menghitung rata-rata setiap kelompok dengan membagi total berat terhadap jumlah data pada kelompok tersebut, termasuk menangani kelompok terakhir jika jumlah datanya tidak penuh.]

### 3. [Soal Latihan Modul 10]
#### soal3.go

```go
package main

import "fmt"

type balita [100]float64

func hitungminmax(bb balita, n int, min *float64, max *float64) {
	*min = bb[0]
	*max = bb[0]

	for i := 1; i < n; i++ {
		if bb[i] < *min {
			*min = bb[i]
		}

		if bb[i] > *max {
			*max = bb[i]
		}
	}
}

func rerata(bb balita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += bb[i]
	}
	return total / float64(n)
}

func main() {
	var bb balita
	var n int
	var min, max float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("masukan berat balita ke - ", i+1, " : ")
		fmt.Scan(&bb[i])
	}
	hitungminmax(bb, n, &min, &max)
	fmt.Printf("Berat balita minimum: %.2f kg \n", min)
	fmt.Printf("Berat balita maximum: %.2f kg \n", max)
	fmt.Printf("Rerata berat balita: %.2f kg", rerata(bb, n))
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_Laprak/blob/main/modul10/output/output-soal3.png)
[Program di atas dibuat untuk mengolah data berat badan balita menggunakan bahasa Go. Pengguna terlebih dahulu memasukkan jumlah data balita, kemudian memasukkan berat masing-masing balita satu per satu ke dalam array bb. Program memiliki fungsi hitungminmax untuk mencari berat badan paling kecil dan paling besar dengan cara membandingkan setiap data dalam array, lalu hasilnya disimpan pada variabel min dan max. Selain itu, terdapat fungsi rerata yang digunakan untuk menghitung rata-rata berat badan balita dengan menjumlahkan seluruh data kemudian membaginya dengan jumlah data. Setelah semua proses selesai, program menampilkan berat minimum, berat maksimum, dan rata-rata berat balita dalam satuan kilogram dengan format dua angka di belakang koma.]