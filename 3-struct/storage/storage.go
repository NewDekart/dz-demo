package storage

import (
	"3-struct/bins"
)

type Storage interface {
	Save(data bins.BinList) error
	Get() (bins.BinList, error)
}
