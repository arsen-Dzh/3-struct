package main

import (
	"crypto/rand"
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

// type BinList struct {
// 	binList []Bin
// }

func newBin(name string) (*Bin, error) {
	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}

	bin := &Bin{
		id:      rand.Text(),
		private: false,
		created: time.Now(),
		name:    name,
	}

	return bin, nil
}

func main() {
	name := promptData("Введите имя bina")
	bin, err := newBin(name)
	if err != nil {
		return
	}
	fmt.Print(bin)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
