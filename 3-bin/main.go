package main

import (
	"fmt"
	"time"
)

func main() {
	bin := newBin("1", true, time.Now(), "Bin")
	fmt.Println(bin)
}

type Bin struct {
	id string
	private bool
	createdAt time.Time
	name string
}

type BinList struct {
	bins []Bin
}

func newBin(id string, private bool, createdAt time.Time, name string) Bin {
	return Bin{
		id: id,
		private: private,
		createdAt: createdAt,
		name: name,
	}
}

func newBinList(bins []Bin) BinList {
	return BinList{
		bins: bins,
	}
}
