package main

import (
	"fmt"
)

func main() {
	usdToEur, usdToRub := inputCurrencyRates()
	eurToRub := usdToRub / usdToEur
	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)
}

func inputCurrencyRates() (float64, float64) {
	fmt.Printf("Введите курс доллара к евро: ")
	var usdToEur float64
	fmt.Scan(&usdToEur)
	fmt.Printf("Введите курс доллара к рублю: ")
	var usdToRub float64
	fmt.Scan(&usdToRub)
	return usdToEur, usdToRub
}

func convert(a, b, c) {

}