package main

import (
	"demo/bin/bins"
	"fmt"
)

func main() {
	id := "123"
	private := true
	name := "Test"

	myBin, err := bins.NewBin(id, private, name)
	if err != nil {
		fmt.Println("Неверный формат id или Name")
		return
	}
	myBin.OutputBin()

	binList := bins.NewBinList()

	binList.AddBin(*myBin)

	binList.OutputBinList()

}
