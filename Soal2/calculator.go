package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run calculator.go [number] [operator] [number]")
		fmt.Println("Supported operators: +, -, *, /")
		return
	}

	// ambil angka dan operator
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid number.")
		return
	}

	operator := os.Args[2]

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid number.")
		return
	}

	// hitung hasil berdasarkan operator
	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: cannot divide by zero.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator. Supported operators: +, -, *, /")
		return
	}

	// tampilkan hasil
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
