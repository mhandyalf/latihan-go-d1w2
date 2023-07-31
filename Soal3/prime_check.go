package main

import (
	"fmt"
)

// fungsi untuk mengecek bilangan prima
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Program untuk memeriksa apakah angka genap atau ganjil, serta apakah angka tersebut bilangan prima.")
	fmt.Print("Masukkan angka: ")

	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid number.")
		return
	}

	// cek apakah angka genap atau ganjil
	if input%2 == 0 {
		fmt.Printf("%d adalah angka Genap.\n", input)
	} else {
		fmt.Printf("%d adalah angka Ganjil. \n", input)
	}

	// cek apakah angka adalah bilangan prima
	if isPrime(input) {
		fmt.Printf("%d adalah bilangan Prima.\n", input)
	} else {
		fmt.Printf("%d bukan bilangan Prima. \n", input)
	}

	// cetak semua bilangan genap dari 1 hingga input
	fmt.Println("Bilangan Genap dari 1 hingga", input, "adalah:")
	for i := 1; i <= input; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
