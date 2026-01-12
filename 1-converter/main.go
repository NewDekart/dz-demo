package main

import "fmt"

const USD_2_RUB = 94
const USD_2_EUR = 1.1
const EUR_2_RUB = USD_2_RUB / USD_2_EUR

func main() {
	fmt.Println("Введите сумму в EUR:")

	eur_sum, err := getSum()

	if err != nil {
		fmt.Println("Ошибка ввода")

		return
	}

	calculate(eur_sum, "EUR", "RUB")

	rub_sum := eur_sum * EUR_2_RUB

	fmt.Printf("Сумма в RUB: %.2f\n", rub_sum)
}

func calculate(value float64, src_currency string, dst_currency string) {}

func getSum() (float64, error) {
	var eur_sum float64

	_, err := fmt.Scan(&eur_sum)

	if err != nil {
		return eur_sum, err
	}

	return eur_sum, nil
}
