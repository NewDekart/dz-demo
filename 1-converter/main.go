package main

import "fmt"

const USD_2_RUB = 94
const USD_2_EUR = 1.1

func main() {
	fmt.Println("Введите сумму в EUR:")

	var eur_sum float64

	fmt.Scan(&eur_sum)

	rub_sum := (eur_sum / USD_2_EUR) * USD_2_RUB

	fmt.Printf("Сумма в RUB: %.2f\n", rub_sum)
}
