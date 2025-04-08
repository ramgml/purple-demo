package main

import (
	"errors"
	"fmt"
)

type RatesMap = map[string]map[string]float64

const USD = "USD"
const EUR = "EUR"
const RUB = "RUB"

const USD_TO_EUR = 0.85
const USD_TO_RUB = 85.0
const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

var rates = RatesMap{
	USD: {
		EUR: USD_TO_EUR,
		RUB: USD_TO_RUB,
	},
	EUR: {
		USD: 1.0 / USD_TO_EUR,
		RUB: EUR_TO_RUB,
	},
	RUB: {
		USD: 1.0 / USD_TO_RUB,
		EUR: 1.0 / EUR_TO_RUB,
	},
}


func main() {
	sourceCurrency := inputSourceCurrency()
	amount := inputAmount()	
	targetCurrency := inputTargetCurrency(sourceCurrency)
	fmt.Printf("Полученные данные: %s %.2f %s\n", sourceCurrency, amount, targetCurrency)
	result, err := convert(&rates, sourceCurrency, amount, targetCurrency)
	if err != nil {
		fmt.Println("Ошибка: Неверный ввод валюты")
		return
	}
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


func convert(rates *RatesMap, sourceCurrency string, amount float64, targetCurrency string) (float64, error) {
	rate := 0.0
	if sourceCurrencyRates, ok := (*rates)[sourceCurrency]; ok {
		if targetCurrencyRate, ok := sourceCurrencyRates[targetCurrency]; ok {
			rate = targetCurrencyRate
		} else {
			return 0, errors.New("unknown target currency")
		}
	} else {
		return 0, errors.New("unknown source currency")
	}
	return amount * rate, nil
}
