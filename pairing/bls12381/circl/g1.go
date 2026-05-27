//nolint:dupl // unavoidable duplication between g1 and g2
package circl

import (
	"crypto/cipher"
	"io"

	bls12381 "github.com/cloudflare/circl/ecc/bls12381"
	"go.dedis.ch/kyber/v4"
)

var _ kyber.SubGroupElement = &G1Elt{}

// G1Elt is a wrapper around a G1 point on the BLS12-381 Circl curve.
type G1Elt struct{ inner bls12381.G1 }

// MarshalBinary returns a compressed point, without any domain separation tag information
func (p *G1Elt) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalBinary populates the point from a compressed point representation.
func (p *G1Elt) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (p *G1Elt) String() string { _ = "STUB: not implemented"; return "" }

func (p *G1Elt) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

// MarshalTo writes a compressed point to the Writer, without any domain separation tag information
func (p *G1Elt) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// UnmarshalFrom populates the point from a compressed point representation read from the Reader.
func (p *G1Elt) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *G1Elt) Equal(p2 kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (p *G1Elt) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *G1Elt) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *G1Elt) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G1Elt) Set(p2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G1Elt) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *G1Elt) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (p *G1Elt) Embed(_ []byte, _ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G1Elt) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *G1Elt) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G1Elt) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G1Elt) Neg(a kyber.Point) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *G1Elt) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *G1Elt) IsInCorrectGroup() bool { _ = "STUB: not implemented"; return false }

var domainG1 = []byte("BLS_SIG_BLS12381G1_XMD:SHA-256_SSWU_RO_NUL_")

func (p *G1Elt) Hash(msg []byte) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }
func (p *G1Elt) Hash2(msg, dst []byte) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}
