/*NIM : 2311102202
NAMA : MUHAMMAD ILHAM SALAM*/

/*package main

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
*/

package main

import "fmt"

func main() {
        var N int
        fmt.Print("Masukkan jumlah peserta: ")
        fmt.Scan(&N)

        var tiket []int
        tiket = make([]int, N)

        for i := 0; i < N; i++ {
                fmt.Printf("Masukkan nomor tiket peserta %d: ", i+1)
                fmt.Scan(&tiket[i])
        }

        var utama, sembako, konsol int

        for _, nomor := range tiket {
                digitSama := true
                digitGenap := true
                temp := nomor
                firstDigit := temp % 10

                for temp > 0 {
                        if temp%10 != firstDigit {
                                digitSama = false
                                break
                        }
                        if temp%10%2 != 0 {
                                digitGenap = false
                        }
                        temp /= 10
                }

                if digitSama && digitGenap {
                        utama++
                        fmt.Println("Hadiah Utama")
                } else if digitSama && !digitGenap {
                        sembako++
                        fmt.Println("Hadiah Sembako")
                } else {
                        konsol++
                        fmt.Println("Hadiah Konsol")
                }
        }

        fmt.Println("Jumlah:")
        fmt.Println("Hadiah Utama:", utama)
        fmt.Println("Hadiah Sembako:", sembako)
        fmt.Println("Hadiah Konsol:", konsol)
}