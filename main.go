package main

import "fmt"

func main() {
	const usdToEur float64 = 0.86
	const usdToRub float64 = 80.87
	const EurToRub float64 = usdToRub / usdToEur
	fmt.Println(EurToRub)

	originCurrency, targetCurrency, sum := getUserInput()
	converter(originCurrency, targetCurrency, sum)
}

func getUserInput() (originCurrency, targetCurrency string, sum float64) {
	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&originCurrency)
	fmt.Print("Введите сумму перевода: ")
	fmt.Scan(&sum)
	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&targetCurrency)
	return originCurrency, targetCurrency, sum
}

func converter(originCurrency, targetCurrency string, sum float64) float64 {
	return 0
}
