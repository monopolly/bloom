package bloom

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dustin/go-humanize"
	"github.com/monopolly/file"
	//testing
	//go test -bench=.
	//go test --timeout 9999999999999s
)

func TestMain(u *testing.T) {
	__(u)

	count := 50000000
	est := 0.00001

	p := New(count, est)
	p = NewBloomer(count)
	p.Load("test.bin")
	select {}

	//p = NewBloomer(count)
	for i := 0; i < count; i++ {
		p.Add([]byte(fmt.Sprint(i)))
	}

	err := p.Save("test.bin")
	if err != nil {
		panic(err)
	}

	err = p.Load("test.bin")
	if err != nil {
		panic(err)
	}

	fmt.Println(humanize.Bytes(uint64(file.Size("test.bin"))))
	select {}
}

func Benchmark1(u *testing.B) {
	u.ReportAllocs()
	u.ResetTimer()
	for n := 0; n < u.N; n++ {

	}
}

func Benchmark2(u *testing.B) {
	u.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

		}
	})
}

func __(u *testing.T) {
	fmt.Printf("\033[1;32m%s\033[0m\n", strings.ReplaceAll(u.Name(), "Test", ""))
}
