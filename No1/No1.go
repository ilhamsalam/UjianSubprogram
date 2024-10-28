/*NIM : 2311102202
NAMA : MUHAMMAD ILHAM SALAM*/

/*package main

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

}*/

package main

import "fmt"

func main() {
        var N int
        fmt.Print("Masukkan bilangan bulat positif lebih dari 100: ")
        fmt.Scan(&N)

        // Hitung jumlah digit
        var count int
        temp := N
        for temp > 0 {
                count++
                temp /= 10
        }

        // Tentukan panjang setiap potongan
        var partLength int
        if count%3 == 0 {
                partLength = count / 3
        } else {
                partLength = count/3 + 1
        }

        // Simpan digit-digit dalam slice
        var digits []int
        for N > 0 {
                digits = append(digits, N%10)
                N /= 10
        }

        // Balik urutan digit (agar sesuai dengan urutan asli bilangan)
        for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
                digits[i], digits[j] = digits[j], digits[i]
        }

        // Potong digit-digit menjadi tiga bagian
        var part1, part2, part3 int
        for i := 0; i < partLength; i++ {
                part1 = part1*10 + digits[i]
        }
        for i := partLength; i < partLength*2; i++ {
                part2 = part2*10 + digits[i]
        }
        for i := partLength * 2; i < count; i++ {
                part3 = part3*10 + digits[i]
        }

        // Hitung rata-rata
        average := int(part1+part2+part3) / 3

        // Tampilkan hasil
        fmt.Println("Hasil pemotongan:")
        fmt.Println(part1, part2, part3)
        fmt.Printf("Rata-rata dari ketiga bilangan:\n%d\n", average)
}