package main

import (
	"bufio"
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
	operantion := inputOperantion()
	numbers := inputNumbers()
	switch operantion {
	case AVG:
		fmt.Printf("Среднее арифметическое: %.2f\n", average(numbers))
	case SUM:
		fmt.Printf("Сумма: %.2f\n", sum(numbers))
	case MED:
		fmt.Printf("Медиана: %.2f\n", median(numbers)) 
	}
}

func inputOperantion() string {
	var operantion string
	for {
		fmt.Println("Введите операцию (AVG, SUM, MED):")
		fmt.Scan(&operantion)
		switch operantion {
		case AVG, SUM, MED:
			return operantion
		default:
			fmt.Println("Неизвестная операция")
			continue
		}
	}
}

func inputNumbers() []float64 {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []float64

	fmt.Println("Введите числа через запятую:")
	scanner.Scan()
	inputString := scanner.Text()

	for _, str := range strings.Split(inputString, ",") {
		if str == "" {
			continue
		}
		number, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
		if err != nil {
			fmt.Println("Ошибка преобразования строки в число:", err)
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func average(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func sum(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func median(numbers []float64) float64 {
	sortedNumbers := sortNumbers(numbers)
	length := len(sortedNumbers)
	medianIndex := length / 2
	if length % 2 == 0 {
		return (sortedNumbers[medianIndex] + sortedNumbers[medianIndex - 1]) / 2
	}
	return sortedNumbers[medianIndex]
}

func sortNumbers(numbers []float64) []float64 {
	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Float64s(sortedNumbers)
	return sortedNumbers
}