package bins

import "time"

type Bin struct {
	id string
	private bool
	createdAt time.Time
	name string
}

type BinList struct {
	bins []Bin
}

func NewBin(id string, private bool, createdAt time.Time, name string) Bin {
	return Bin{
		id: id,
		private: private,
		createdAt: createdAt,
		name: name,
	}
}

func NewBinList(bins []Bin) BinList {
	return BinList{
		bins: bins,
	}
}