package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const AVG = "AVG"
const SUM = "SUM"
const MED = "MED"

func main() {
	operantion := inputOperation()
	numbers := inputNumbers()
	switch operantion {
	case AVG:
		avg, err := average(numbers)
		if err != nil {
			fmt.Println("Ошибка: Неверный ввод чисел")
			return
		}
		fmt.Printf("Среднее арифметическое: %.2f\n", avg)
	case SUM:
		sum, err := sum(numbers)
		if err != nil {
			fmt.Println("Ошибка: Неверный ввод чисел")
			return
		}
		fmt.Printf("Сумма: %.2f\n", sum)
	case MED:
		med, err := median(numbers)
		if err != nil {
			fmt.Println("Ошибка: Неверный ввод чисел")
			return
		}
		fmt.Printf("Медиана: %.2f\n", med) 
	}
}

func inputOperation() string {
	var operation string
	for {
		fmt.Println("Введите операцию (AVG, SUM, MED):")
		fmt.Scan(&operation)
		switch operation {
		case AVG, SUM, MED:
			return operation
		default:
			fmt.Println("Неизвестная операция")
			continue
		}
	}
}

func inputNumbers() []float64 {
	var numbers []float64
	mainLoop:
		for {
			fmt.Println("Введите числа через запятую:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			inputString := scanner.Text()
			for _, str := range strings.Split(inputString, ",") {
				if str == "" {
					continue
				}
				number, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
				if err != nil {
					fmt.Println("Ошибка преобразования строки в число:", str)
					continue mainLoop
				}
				numbers = append(numbers, number)
			}
			return numbers
		} 
}

func average(numbers []float64) (float64, error) {
	var sum float64
	if len(numbers) == 0 {
		return 0, errors.New("no numbers provided")
	}
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers)), nil
}

func sum(numbers []float64) (float64, error) {
	var sum float64
	if len(numbers) == 0 {
		return 0, errors.New("no numbers provided")
	}
	for _, number := range numbers {
		sum += number
	}
	return sum, nil
}

func median(numbers []float64) (float64, error) {
	length := len(numbers)
	if length == 0 {
		return 0, errors.New("no numbers provided")
	}
	sortedNumbers := sortNumbers(numbers)
	medianIndex := length / 2
	if length % 2 == 0 {
		return (sortedNumbers[medianIndex] + sortedNumbers[medianIndex - 1]) / 2, nil
	}
	return sortedNumbers[medianIndex], nil
}

func sortNumbers(numbers []float64) []float64 {
	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Float64s(sortedNumbers)
	return sortedNumbers
}