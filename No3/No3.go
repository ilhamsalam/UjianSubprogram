/*package main

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
}*/

package main

import (
    "fmt"
)

func pembagian(n, m int) int {
    if n < m {
        return 0
    }
    return 1 + pembagian(n-m, m)
}

func modulus(n, m int) int {
    if n < m {
        return n
    }
    return modulus(n-m, m)
}

func main() {
    var n, m int
    fmt.Print("Masukkan bilangan pembilang (n): ")
    fmt.Scan(&n)
    fmt.Print("Masukkan bilangan penyebut (m): ")
    fmt.Scan(&m)
    fmt.Printf("Hasil dari pembagian %d dengan %d adalah: %d\n", n, m, pembagian(n, m))
    fmt.Printf("Sisa dari pembagian adalah: %d\n", modulus(n, m))
}
