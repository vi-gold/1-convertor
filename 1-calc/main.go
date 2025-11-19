package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation, listNum := getInputUser()
	arrayNum := convertStringToArray(listNum)
	resultCalc := calculateOperation(arrayNum, operation)
	fmt.Printf("Операция %s с числами %s равна %.2f", operation, listNum, resultCalc)
}

func getInputUser() (operation, listNum string) {
	for {
		fmt.Print("Введите операцию (avg, sum, med): ")
		fmt.Scan(&operation)
		resultCheck, err := checkUserInputOperation(operation)
		if !resultCheck {
			fmt.Println(err)
			continue
		}
		break
	}
	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&listNum)
	return operation, listNum

}

func checkUserInputOperation(operation string) (bool, error) {
	if operation == "avg" || operation == "sum" || operation == "med" {
		return true, nil
	} else {
		return false, errors.New("ОШИБКА: ВВЕДЕНЫ НЕКОРРЕКТНЫЕ ДАННЫЕ")
	}
}

func convertStringToArray(listNum string) []float64 {
	arrayStrings := strings.Split(listNum, ",")
	arrayNum := make([]float64, 0, len(arrayStrings))
	for _, value := range arrayStrings {
		convertedNum, err := strconv.ParseFloat(value, 64)
		if err == nil {
			arrayNum = append(arrayNum, convertedNum)
		}
	}
	return arrayNum
}

func calculateOperation(arrayNum []float64, operation string) (result float64) {
	if operation == "avg" {
		for _, value := range arrayNum {
			result += value
		}
		result = result / float64(len(arrayNum))
	}
	if operation == "sum" {
		for _, value := range arrayNum {
			result += value
		}
	}
	if operation == "med" {
		sort.Float64s(arrayNum)
		if len(arrayNum)%2 == 1 {
			result = arrayNum[len(arrayNum)/2]
		} else {
			median1 := arrayNum[len(arrayNum)-1]
			median2 := arrayNum[len(arrayNum)]
			result = (median1 + median2) / 2
		}
	}
	return result
}
