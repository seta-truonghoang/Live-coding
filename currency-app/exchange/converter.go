package exchange

import "fmt"

var rates map[string]float64

func init() {
	rates = map[string]float64{
		"USD": 1.0,
		"VND": 25400.0,
		"EUR": 0.92,
	}
	fmt.Println("Exchange package initialized!")
}

const Version = "1.0.0"

func Convert(amount float64, from, to string) (float64, error) {
	fromRate, ok1 := rates[from]
	toRate, ok2 := rates[to]

	if !ok1 || !ok2 {
		return 0, fmt.Errorf("currency not supported")
	}

	result := (amount / fromRate) * toRate
	return result, nil
}
