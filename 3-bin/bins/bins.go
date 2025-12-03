package bins

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	CreatedAt time.Time
	name      string
}

func (newBin *Bin) OutputBin() {
	fmt.Println(newBin.id, newBin.private, newBin.CreatedAt, newBin.name)
}

func NewBin(id string, private bool, name string) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("EMPTY_ID_OR_NAME")
	}
	newBin := &Bin{
		id:        id,
		private:   private,
		CreatedAt: time.Now(),
		name:      name,
	}
	return newBin, nil
}

type BinList struct {
	items []Bin
}

func (binList *BinList) OutputBinList() {
	fmt.Println(binList.items)
}

func (binList *BinList) AddBin(myBin Bin) {
	binList.items = append(binList.items, myBin)
}

func NewBinList() *BinList {
	newBinList := &BinList{
		items: []Bin{},
	}
	return newBinList
}
