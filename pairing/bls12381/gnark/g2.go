//go:build !constantTime

package gnark

import (
	"crypto/cipher"
	"io"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"go.dedis.ch/kyber/v4"
)

var _ kyber.SubGroupElement = &G2Elt{}

// G2Elt is a wrapper around the Gnark G2 point type.
type G2Elt struct{ inner bls12381.G2Jac }

// MarshalBinary returns a compressed point, without any domain separation tag information
func (p *G2Elt) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalBinary populates the point from a compressed point representation.
func (p *G2Elt) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (p *G2Elt) String() string { _ = "STUB: not implemented"; return "" }

func (p *G2Elt) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

// MarshalTo writes a compressed point to the Writer, without any domain separation tag information
func (p *G2Elt) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// UnmarshalFrom populates the point from a compressed point representation read from the Reader.
func (p *G2Elt) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *G2Elt) Equal(p2 kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (p *G2Elt) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *G2Elt) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *G2Elt) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G2Elt) Set(p2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G2Elt) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *G2Elt) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (p *G2Elt) Embed(_ []byte, _ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G2Elt) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *G2Elt) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G2Elt) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G2Elt) Neg(a kyber.Point) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *G2Elt) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G2Elt) IsInCorrectGroup() bool { _ = "STUB: not implemented"; return false }

var domainG2 = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

func (p *G2Elt) Hash(msg []byte) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }
func (p *G2Elt) Hash2(msg, dst []byte) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}
