//go:build !constantTime

package bn256

import (
	"crypto/cipher"
	"errors"
	"io"
	"math/big"

	"go.dedis.ch/kyber/v4"
)

var marshalPointID1 = [8]byte{'b', 'n', '2', '5', '6', '.', 'g', '1'}
var marshalPointID2 = [8]byte{'b', 'n', '2', '5', '6', '.', 'g', '2'}
var marshalPointIDT = [8]byte{'b', 'n', '2', '5', '6', '.', 'g', 't'}

var ErrTypeCast = errors.New("invalid type cast")

type pointG1 struct {
	g *curvePoint
}

func newPointG1() *pointG1 { _ = "STUB: not implemented"; return nil }

func (p *pointG1) Equal(q kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (p *pointG1) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG1) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG1) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG1) Set(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Clone makes a hard copy of the point
func (p *pointG1) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG1) EmbedLen() int {
	_ = "STUB: not implemented"
	// 2^255 is ~size of the curve P
	// minus one byte for randomness
	// minus one byte for len(data)
	return 0
}

func (p *pointG1) Embed(data []byte, rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	// How many bytes to embed?
	return *new(kyber.Point)
}

// Pick a random point, with optional embedded data

// Encode length in low 8 bits
// Copy in data to embed

func (p *pointG1) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// extract length byte

func (p *pointG1) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// p = a + b

func (p *pointG1) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG1) Neg(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG1) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG1) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	// Clone is required as we change the point
	return nil, nil
}

// Take a copy so that p is not written to, so calls to MarshalBinary
// are threadsafe.

func (p *pointG1) MarshalID() [8]byte { _ = "STUB: not implemented"; return nil }

func (p *pointG1) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointG1) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

// This is the point at infinity

func (p *pointG1) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointG1) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG1) ElementSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG1) String() string { _ = "STUB: not implemented"; return "" }

func (p *pointG1) Hash(m []byte) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

// hashes a byte slice into a curve point represented by two big.Int's
// ideally we want to do this using gfP, but gfP doesn't have a ModSqrt function
func hashToPoint(m []byte) (*big.Int, *big.Int) { _ = "STUB: not implemented"; return nil, nil }

func deriveY(x *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

type pointG2 struct {
	g *twistPoint
}

func newPointG2() *pointG2 { _ = "STUB: not implemented"; return nil }

func (p *pointG2) Equal(q kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (p *pointG2) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG2) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG2) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG2) Set(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Clone makes a hard copy of the field
func (p *pointG2) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointG2) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG2) Embed(_ []byte, _ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG2) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pointG2) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// p = a + b

func (p *pointG2) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG2) Neg(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG2) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointG2) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	// Clone is required as we change the point during the operation
	return nil, nil
}

func (p *pointG2) MarshalID() [8]byte { _ = "STUB: not implemented"; return nil }

func (p *pointG2) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointG2) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

// This is the point at infinity.

func (p *pointG2) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointG2) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG2) ElementSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointG2) String() string { _ = "STUB: not implemented"; return "" }

type pointGT struct {
	g *gfP12
}

func newPointGT() *pointGT { _ = "STUB: not implemented"; return nil }

func (p *pointGT) Equal(q kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (p *pointGT) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointGT) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointGT) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Set(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Clone makes a hard copy of the point
func (p *pointGT) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointGT) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (p *pointGT) Embed(_ []byte, _ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pointGT) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Neg(q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pointGT) MarshalID() [8]byte { _ = "STUB: not implemented"; return nil }

func (p *pointGT) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointGT) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

// TODO: check if point is on curve

func (p *pointGT) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pointGT) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointGT) ElementSize() int { _ = "STUB: not implemented"; return 0 }

func (p *pointGT) String() string { _ = "STUB: not implemented"; return "" }

func (p *pointGT) Finalize() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *pointGT) Miller(p1, p2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *pointGT) Pair(p1, p2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}
