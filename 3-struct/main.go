package main

import (
	"3-struct/bins"
	"3-struct/storage"
	"fmt"
)

func main() {
	bin := bins.NewBin("id123", true, "Vasya")
	binList := bins.NewBinList()

	binList = append(binList, bin)

	storage := storage.JsonStorage{}

	storage.Save(binList)

	result, _ := storage.Get()

	fmt.Println(result)
}
