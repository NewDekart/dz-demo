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
	operation, err := getOperation()

	if err != nil {
		showError(err)
		return
	}

	numbers, err := getNumbers()

	if err != nil {
		showError(err)
		return
	}

	result := calculate(operation, numbers)

	showResult(result)
}

func getOperation() (string, error) {
	var operation string

	fmt.Println("Введите операцию (AVG, SUM, MED):")
	operation = scan()

	if operation != AVG && operation != SUM && operation != MED {
		return "", errors.New("Неизвестная операция")
	}

	return operation, nil
}

func getNumbers() ([]float64, error) {
	fmt.Println("Введите числа через запятую:")
	numbersString := scan()

	return parseNumbers(numbersString)
}

func parseNumbers(numbersString string) ([]float64, error) {
	parts := strings.Split(numbersString, ",")

	numbers := make([]float64, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		number, err := strconv.ParseFloat(trimmed, 64)

		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	if len(numbers) == 0 {
		return nil, errors.New("Нет чисел")
	}

	return numbers, nil
}

func calculate(operation string, numbers []float64) float64 {
	switch operation {
	case AVG:
		return calculateAvg(numbers)
	case SUM:
		return calculateSum(numbers)
	case MED:
		return calculateMed(numbers)
	default:
		panic("Unknown operation")
	}
}

func calculateAvg(numbers []float64) float64 {
	sum := calculateSum(numbers)
	return sum / float64(len(numbers))
}

func calculateSum(numbers []float64) float64 {
	var sum float64

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func calculateMed(numbers []float64) float64 {
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)

	sort.Float64s(sorted)

	n := len(sorted)

	if n%2 == 1 {
		return sorted[n/2]
	}

	return (sorted[n/2-1] + sorted[n/2]) / 2
}

func showResult(result float64) {
	fmt.Println("Результат:")
	fmt.Println(result)
}

func showError(err error) {
	fmt.Println(err)
}

func scan() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}
