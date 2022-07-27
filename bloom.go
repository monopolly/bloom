package bloom

import (
	"os"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/monopolly/file"
)

type Bloomer struct {
	count uint
	db    *bloom.BloomFilter
}

// read balance
func New(count int, esimate float64) (a *Bloomer) {
	a = new(Bloomer)
	a.count = uint(count)
	a.db = bloom.NewWithEstimates(a.count, esimate)
	return
}

// read balance
func NewBloomer(count int) (a *Bloomer) {
	a = new(Bloomer)
	a.count = uint(count)
	a.db = bloom.NewWithEstimates(a.count, 0.00000000001)
	return
}

// set medium false positive percent
func (a *Bloomer) Medium() {
	a.db = bloom.NewWithEstimates(a.count, 0.0000001)
}

// set false positive percent
func (a *Bloomer) Set(percent float64) {
	a.db = bloom.NewWithEstimates(a.count, percent)
}

func (a *Bloomer) Add(b []byte) {
	a.db.Add(b)
}
func (a *Bloomer) Has(b []byte) bool {
	return a.db.Test(b)
}

func (a *Bloomer) Save(fn string) (err error) {
	file.Delete(fn)
	f, err := os.Create(fn)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = a.db.WriteTo(f)
	return
}

func (a *Bloomer) Load(fn string) (err error) {
	f, err := os.Open(fn)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = a.db.ReadFrom(f)
	return
}
