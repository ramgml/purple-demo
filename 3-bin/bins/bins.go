package bins

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func NewBin(id string, private bool, createdAt time.Time, name string) Bin {
	return Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func (binList *BinList) AddBin (bin *Bin) {
	binList.Bins = append(binList.Bins, *bin)
}


func NewBinList(bins []Bin) BinList {
	return BinList{
		Bins: bins,
	}
}

