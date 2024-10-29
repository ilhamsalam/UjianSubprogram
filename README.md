# Ujian Sub Program (Tipe C)
Penjelasan singkat pada tiap soal ujian sub program tipe c

## Setiap soal
Pada setiap soal saya menggunakan beberapa komponen :
- `Package main` agar program dapat dieksekusi
- `import fmt` agar dapat menggunakan beberapa operasi dasar bahasa program `go`
- `main()` sebagai titik awal program dan dieksekusi ketika program dijalankan

## Soal 1 (versi ujian)
```go
package main

import "fmt"

func pemotongan(N int) {
        if N < 10 {
                fmt.Print(N)
                return
        }

        pemotongan(N / 10)

        fmt.Print(" ")

        fmt.Print(N % 10)

}

func main() {
        var N int
        fmt.Scan(&N)
        pemotongan(N)
        fmt.Println()

}
```
Operasi diatas merupakan operasi pemisah antar bilangan dalam satu kesatuan

## Soal 2 (versi ujian)
```go
package main

import "fmt"

func main() {
    var angka [5] int
	var n int

	fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Scan(&angka[i])
    }

    for i := 0; i < 5; i++ {
		if angka[i] == angka[i+1] && angka[i] %2 == 0{
			fmt.Printf("Hadiah utama")
		}
        
    }
}
```
Operasi diatas merupakan operasi menentukan hadiah yang didapatkan peserta

## Soal 3 (versi ujian)
```go
package main

import "fmt"

func pembagian(n, m int)int{
    return n/m
}

func modulus(n, m int)int{
	return n%m
}

func main() {
    var n, m int
    fmt.Scan(&n, &m)
	fmt.Println(pembagian(n, m))
	fmt.Println(modulus(n, m))
}
```

Operasi diatas merupakan operasi  untuk menghitung pembagian dan sisa bagi dari dua  buah bilangan
