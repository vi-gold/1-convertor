package main

import "fmt"

func main() {
	const usdToEur float64 = 0.86
	const usdToRub float64 = 80.87
	const EurToRub float64 = usdToRub / usdToEur
	fmt.Print(EurToRub)
}
