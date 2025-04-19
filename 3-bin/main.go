package main

import (
	"fmt"
	"time"
	"3-bin/bins"
	"3-bin/storage"
)

func main() {
	bin := bins.NewBin("1", true, time.Now(), "Bin")
	var binList bins.BinList
	binList.AddBin(&bin)
	dataStorage := storage.NewStorage(&binList)
	err := dataStorage.Save()
	if err != nil {
		fmt.Println("Не удалось сохранить файл")
	}
	loaded, _ := dataStorage.List()
	fmt.Println(loaded)
	fmt.Println(bin)
}
