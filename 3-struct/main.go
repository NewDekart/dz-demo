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

	var s storage.Storage = &storage.JsonStorage{}

	s.Save(binList)

	result, _ := s.Get()

	fmt.Println(result)
}
