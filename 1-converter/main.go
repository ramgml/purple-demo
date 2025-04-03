package main

import (
	"errors"
	"fmt"
)

const USD = "USD"
const EUR = "EUR"
const RUB = "RUB"

const USD_TO_EUR = 0.85
const USD_TO_RUB = 85.0
const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR


func main() {
	sourceCurrency := inputSourceCurrency()
	amount := inputAmount()	
	targetCurrency := inputTargetCurrency(sourceCurrency)
	fmt.Printf("Полученные данные: %s %.2f %s\n", sourceCurrency, amount, targetCurrency)
	result := convert(sourceCurrency, amount, targetCurrency)
	fmt.Printf("%.2f %s = %.2f %s\n", amount, sourceCurrency, result, targetCurrency)
}

func inputSourceCurrency() string {
	var sourceCurrency string
	for {
		fmt.Println("Введите исходную валюту:")
		fmt.Scan(&sourceCurrency)
		switch sourceCurrency {
		case USD, EUR, RUB:
			return sourceCurrency
		default:
			fmt.Println("Неизвестная валюта")
			continue
		}
	}
}

func inputAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите количество валюты:\n")
		_, err := fmt.Scan(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Количество валюты должно быть числом больше нуля")
			//Очистка буфера ввода
			discard := ""
			fmt.Scan(&discard)

			continue
		}
		return amount
	}
}

func inputTargetCurrency(sourceCurrency string) string {
	var targetCurrency string

	for {
		fmt.Printf("Введите целевую валюту: \n")
		fmt.Scan(&targetCurrency)
		switch targetCurrency {
		case USD, EUR, RUB:
			if targetCurrency == sourceCurrency {
				fmt.Println("Целевая валюта должна отличаться от исходной")
				continue
			}
		default:
			fmt.Println("Неизвестная валюта")
			continue
		}
		return targetCurrency
	}
}

func convert(sourceCurrency string, amount float64, targetCurrency string) float64{
	switch sourceCurrency {
	case USD:
		switch targetCurrency {
		case EUR:
			return amount * USD_TO_EUR
		case RUB:
			return amount * USD_TO_RUB
		}
	case EUR:
		switch targetCurrency {
		case USD:
			return amount / USD_TO_EUR
		case RUB:
			return amount * EUR_TO_RUB
		}
	case RUB:
		switch targetCurrency {
		case USD:
			return amount / USD_TO_RUB
		case EUR:
			return amount / EUR_TO_RUB
		}
	}
	return 0
}
