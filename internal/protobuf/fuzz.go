// To fuzz-test this package:
//
//   $ go get -u github.com/dvyukov/go-fuzz/go-fuzz
//   $ go get -u github.com/dvyukov/go-fuzz/go-fuzz-build
//   $ cd `go env GOPATH`
//   $ go-fuzz-build go.dedis.ch/protobuf
//   $ go-fuzz -workdir=workdir -bin protobuf-fuzz.zip

//go:build gofuzz
// +build gofuzz

package protobuf

type t1 [32]byte
type t2 struct {
	X, Y t1
	Sl   []bool
	T3   t3
	T3s  [3]t3
}
type t3 struct {
	I int
	F float64
	B bool
}

func Fuzz(data []byte) int { _ = "STUB: not implemented"; return 0 }
