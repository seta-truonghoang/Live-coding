package main

import (
	"currency-app/exchange" // Import local package
	"fmt"
)

func main() {
	amount := 100.0
	from := "USD"
	to := "VND"

	fmt.Printf("Library Version: %s\n", exchange.Version)

	result, err := exchange.Convert(amount, from, to)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
}
