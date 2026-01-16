package main

import "time"

type Bin struct {
	id        string
	private   bool
	createAdt time.Time
	name      string
}

type BinList []Bin

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		id:        id,
		private:   private,
		createAdt: time.Now(),
		name:      name,
	}
}

func NewBinList() BinList {
	return BinList{}
}
