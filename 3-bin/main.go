package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func (newBin *Bin) outputBin() {
	fmt.Println(newBin.id, newBin.private, newBin.createdAt, newBin.name)
}

func (newBin *Bin) outputBinList(binList []Bin) {
	fmt.Println(binList)
}

func newBin(id string, private bool, name string) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("EMPTY_ID_OR_NAME")
	}
	newBin := &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
	return newBin, nil
}

func main() {
	id := "123"
	private := true
	name := "Test"

	myBin, err := newBin(id, private, name)
	if err != nil {
		fmt.Println("Неверный формат id или Name")
		return
	}
	myBin.outputBin()
	binList := []Bin{}
	binList = append(binList, *myBin)
	myBin.outputBinList(binList)

}
