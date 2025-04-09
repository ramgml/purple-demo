package main

import (
	"fmt"
	"time"
	"3-bin/bins"
)

func main() {
	bin := bins.NewBin("1", true, time.Now(), "Bin")
	fmt.Println(bin)
}
