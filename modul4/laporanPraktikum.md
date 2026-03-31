# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">[Farel Juliyandra Restu Hermawan] - [109082530038]</p>

## Unguided 

### 1. [Soal 1]
#### soal1.go

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}
func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}
func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var p1, c1, p2, c2 int
	permutation(a, c, &p1)
	combination(a, c, &c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_4/blob/main/modul4/output/output-soal1.png)
[Program ini menghitung permutasi dari dua pasangan bilangan input (a, c) dan (b, d). Fungsi factorial digunakan untuk menghitung faktorial suatu bilangan dengan menyimpan hasilnya melalui pointer (*int). Fungsi permutation menghitung ( P(n,r) = n!/(n-r)! ), sedangkan combination menghitung ( C(n,r) = n!/(r!(n-r)!) ), keduanya juga menggunakan hasil dari factorial. Di main, program membaca empat angka, lalu menghitung permutasi dan kombinasi untuk masing-masing pasangan dan mencetak hasilnya dalam dua baris, masing-masing berisi nilai permutasi dan kombinasi.]

### 2. [Soal 2]
#### soal2.go

```go
package main

import (
	"fmt"
)

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0

	var waktu int
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal += 1
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var maxSoal, minWaktu int
	minWaktu = 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		var soal, skor int
		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			maxSoal = soal
			minWaktu = skor
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_4/blob/main/modul4/output/output-soal2.png)
[program ini digunakan untuk menentukan pemenang kompetisi berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu tercepat. Prosedur hitungSkor membaca waktu pengerjaan 8 soal untuk setiap peserta, lalu menghitung berapa soal yang selesai (waktu ≤ 300 menit) dan menjumlahkan total waktunya menggunakan pointer agar hasil langsung tersimpan ke variabel luar. Di main, program membaca nama peserta berulang hingga “Selesai”, kemudian memanggil prosedur tersebut untuk setiap peserta dan membandingkan hasilnya: peserta dengan jumlah soal terbanyak menjadi pemenang, dan jika sama, dipilih yang memiliki total waktu lebih kecil. Akhirnya, nama pemenang, jumlah soal, dan total waktu ditampilkan.]

### 3. [Soal 3]
#### soal3.go

```go
package main

import "fmt"

func collatz(n int) {
	for n != 1 {
		fmt.Print(n, " ")

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var n int
	fmt.Scan(&n)

	collatz(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/faarrreeeelll/Farel_Juliyandra_Restu-Hermawan_109082530038_Modul_4/blob/main/modul4/output/output-soal3.png)
[Program ini mencetak deret Collatz dari sebuah bilangan input n hingga mencapai 1. Prosedur collatz akan melakukan perulangan selama n belum sama dengan 1, setiap langkah mencetak nilai n, lalu jika n genap dibagi 2, sedangkan jika ganjil dihitung dengan rumus 3n + 1. Proses ini terus berulang sampai akhirnya mencapai 1, lalu angka 1 dicetak sebagai akhir deret.]



