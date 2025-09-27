package bins

import (
	"errors"
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

func NewBin(name string) (*Bin, error) {
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

func NewBinList(bin *Bin) *BinList {

	binArray := []Bin{*bin}

	binList := &BinList{
		bins: binArray,
	}

	return binList
}
