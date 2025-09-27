package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id      string
	private bool
	created time.Time
	name    string
}

type BinList struct {
	bins []Bin
}

func newBin(name string) (*Bin, error) {
	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}

	bin := &Bin{
		id:      "default_id",
		private: false,
		created: time.Now(),
		name:    name,
	}

	return bin, nil
}

func newBinList(bin *Bin) *BinList {

	binArray := []Bin{*bin}

	binList := &BinList{
		bins: binArray,
	}

	return binList
}

func main() {
	for {
		var isCreated string
		fmt.Println("Создать bin yes/no?")
		fmt.Scan(&isCreated)
		if isCreated != "yes" {
			break
		}
		name := promptData("Введите имя bina")
		bin, err := newBin(name)
		if err != nil {
			return
		}
		fmt.Print(bin)
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
