package main

import (
	"errors"
	"fmt"
)

type usdMap = map[string]float64
type eurMap = map[string]float64
type rubMap = map[string]float64
type testMap = map[string]map[string]float64

func main() {

	usd := usdMap{"eur": 0.86, "rub": 80.87}
	eur := eurMap{"rub": usd["rub"] / usd["eur"], "usd": usd["rub"] / usd["eur"] / usd["rub"]}
	rub := rubMap{"usd": 1 / usd["rub"], "eur": 1 / eur["rub"]}
	conv := testMap{"usd": usd, "eur": eur, "rub": rub}

	// const usdToEur float64 = 0.86
	// const usdToRub float64 = 80.87
	// const eurToRub float64 = usdToRub / usdToEur
	// const eurToUsd float64 = eurToRub / usdToRub
	// const rubToUsd float64 = 1 / usdToRub
	// const rubToEur float64 = 1 / eurToRub
	originCurrency, targetCurrency, sum := getUserInput()
	result := convert(originCurrency, targetCurrency, sum, &conv)
	fmt.Printf("Конвертация суммы %.2f%s в %s равна %.2f\n",
		sum, originCurrency, targetCurrency, result)
}

func getUserInput() (originCurrency, targetCurrency string, sum float64) {
	currencies := "(usd, rub, eur): "
	for {
		fmt.Print("Введите исходную валюту " + currencies)
		fmt.Scan(&originCurrency)
		resultCheck, err := checkUserInputOriginCurrency(originCurrency)
		if !resultCheck {
			fmt.Println(err)
			continue
		}
		break
	}
	currencies = getAvailableCurrencies(originCurrency)
	for {
		fmt.Print("Введите сумму перевода: ")
		fmt.Scan(&sum)
		resultCheck, err := checkUserInputSum(sum)
		if !resultCheck {
			fmt.Println(err)
			continue
		}
		break
	}
	for {
		fmt.Print("Введите целевую валюту " + currencies)
		fmt.Scan(&targetCurrency)
		resultCheck, err := checkUserInputTargetCurrency(targetCurrency, originCurrency)
		if !resultCheck {
			fmt.Println(err)
			continue
		}
		break
	}
	return originCurrency, targetCurrency, sum
}

func checkUserInputOriginCurrency(originCurrency string) (bool, error) {
	if originCurrency == "rub" || originCurrency == "usd" || originCurrency == "eur" {
		return true, nil
	} else {
		return false, errors.New("ОШИБКА: ВВЕДЕНЫ НЕКОРРЕКТНЫЕ ДАННЫЕ")
	}
}

func checkUserInputSum(sum float64) (bool, error) {
	if sum > 0 {
		return true, nil
	} else {
		return false, errors.New("ОШИБКА: ВВЕДЕНЫ НЕКОРРЕКТНЫЕ ДАННЫЕ")
	}
}

func checkUserInputTargetCurrency(targetCurrency, originCurrency string) (bool, error) {
	if targetCurrency == "rub" || targetCurrency == "usd" || targetCurrency == "eur" {
		if targetCurrency == originCurrency {
			return false, errors.New("ОШИБКА: ИСХОДНАЯ И ЦЕЛЕВАЯ ВАЛЮТА СОВПАДАЮТ")
		}
		return true, nil
	} else {
		return false, errors.New("ОШИБКА: ВВЕДЕНЫ НЕКОРРЕКТНЫЕ ДАННЫЕ")
	}
}

func getAvailableCurrencies(originCurrency string) (availableCurrencies string) {
	switch {
	case originCurrency == "usd":
		availableCurrencies = "(rub, eur): "
	case originCurrency == "rub":
		availableCurrencies = "(usd, eur): "
	case originCurrency == "eur":
		availableCurrencies = "(rub, usd): "
	}
	return availableCurrencies
}

func convert(originCurrency, targetCurrency string, sum float64, conv *testMap) (result float64) {
	result = (*conv)[originCurrency][targetCurrency] * sum
	return result
	// switch {
	// case originCurrency == "usd":
	// 	switch {
	// 	case targetCurrency == "rub":
	// 		result = usdToRub * sum
	// 	case targetCurrency == "eur":
	// 		result = usdToEur * sum
	// 	}
	// case originCurrency == "eur":
	// 	switch {
	// 	case targetCurrency == "usd":
	// 		result = eurToUsd * sum
	// 	case targetCurrency == "rub":
	// 		result = eurToRub * sum
	// 	}
	// case originCurrency == "rub":
	// 	switch {
	// 	case targetCurrency == "usd":
	// 		result = rubToUsd * sum
	// 	case targetCurrency == "eur":
	// 		result = rubToEur * sum
	// 	}
	// }
}
