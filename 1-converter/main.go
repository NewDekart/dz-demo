package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const FORMAT_COURSE_STRING = "%s_%s"

const USD_2_RUB = 94
const USD_2_EUR = 1.1
const EUR_2_RUB = USD_2_RUB / USD_2_EUR

type Currency string

const (
	EUR Currency = "EUR"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

func main() {
	var step int
	var dst, src Currency
	var sum float64

	m := map[string]float64{
		fmt.Sprintf(FORMAT_COURSE_STRING, RUB, USD): 1 / USD_2_RUB,
		fmt.Sprintf(FORMAT_COURSE_STRING, RUB, EUR): 1 / EUR_2_RUB,
		fmt.Sprintf(FORMAT_COURSE_STRING, USD, EUR): USD_2_EUR,
		fmt.Sprintf(FORMAT_COURSE_STRING, USD, RUB): USD_2_RUB,
		fmt.Sprintf(FORMAT_COURSE_STRING, EUR, RUB): EUR_2_RUB,
		fmt.Sprintf(FORMAT_COURSE_STRING, EUR, USD): 1 / USD_2_EUR,
	}

	for {
		if step == 0 {
			value, err := getCurrency("")

			if err != nil {
				fmt.Println(err)

				continue
			}

			src = value

			step++
		}

		if step == 1 {
			value, err := getSum()

			if err != nil {
				fmt.Println(err)

				continue
			}

			sum = value

			step++
		}

		if step == 2 {
			value, err := getCurrency(src)

			if err != nil {
				fmt.Println(err)

				continue
			}

			dst = value

			break
		}
	}

	result, err := calculate(sum, src, dst, &m)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("Сумма в %s: %.2f\n", dst, result)
}

func calculate(value float64, src_currency Currency, dst_currency Currency, m *map[string]float64) (float64, error) {
	var rate float64

	course := fmt.Sprintf(FORMAT_COURSE_STRING, src_currency, dst_currency)

	rate, ok := (*m)[course]

	if !ok {
		return rate, errors.New("Отствуют данные о курсе валют")
	}

	return value * rate, nil
}

func getSum() (float64, error) {
	fmt.Println("Введите сумму:")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	sum, err := strconv.ParseFloat(input.Text(), 64)

	if err != nil {
		return sum, errors.New("Неверный формат суммы")
	}

	if sum <= 0 {
		return sum, errors.New("Сумма должна быть больше 0")
	}

	return sum, nil
}

func getCurrency(src Currency) (Currency, error) {
	var currency_description string

	if src == "" {
		currency_description = "исходную"
	} else {
		currency_description = "желаемую"
	}

	allTips := []Currency{EUR, RUB, USD}

	if src != "" {
		newTips := []Currency{}

		for _, value := range allTips {
			if value != src {
				newTips = append(newTips, value)
			}
		}

		allTips = newTips
	}

	fmt.Printf("Введите %s валюту(%v):\n", currency_description, allTips)

	var result Currency

	_, err := fmt.Scan(&result)

	if err != nil {
		return "", err
	}

	if result != EUR && result != RUB && result != USD {
		return "", errors.New("Неизвестная валюта")
	}

	if result == src {
		return "", errors.New("Исходная и конечная валюты совпадают")
	}

	return result, nil
}
