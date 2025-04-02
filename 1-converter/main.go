package main

import (
	"fmt"
)

const USD_TO_EUR float64 = 0.85
const USD_TO_RUB float64 = 100


func main() {
	eur_to_rub := USD_TO_RUB / USD_TO_EUR
	fmt.Printf("1 EUR = %.2f RUB\n", eur_to_rub)
}