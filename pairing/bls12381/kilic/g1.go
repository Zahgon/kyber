//go:build !constantTime

package kilic

import (
	"crypto/cipher"
	"io"

	bls12381 "github.com/kilic/bls12-381"
	"go.dedis.ch/kyber/v4"
)

// domainG1 is the DST used for hash to curve on G1, this is the default from the RFC.
var domainG1 = []byte("BLS_SIG_BLS12381G1_XMD:SHA-256_SSWU_RO_NUL_")

func DefaultDomainG1() []byte {
	_ = "STUB: not implemented"

	// G1Elt is a kyber.Point holding a G1 point on BLS12-381 curve
	return nil
}

type G1Elt struct {
	p *bls12381.PointG1
	// domain separation tag. We treat a 0 len dst as the default value as per the RFC "Tags MUST have nonzero length"
	dst []byte

	kyber.Point
	kyber.HashablePoint
}

func NullG1(dst ...byte) *G1Elt { _ = "STUB: not implemented"; return nil }

func newG1(p *bls12381.PointG1, dst []byte) *G1Elt { _ = "STUB: not implemented"; return nil }

func (k *G1Elt) Equal(k2 kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (k *G1Elt) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (k *G1Elt) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (k *G1Elt) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (k *G1Elt) Set(q kyber.Point) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (k *G1Elt) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (k *G1Elt) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (k *G1Elt) Embed(_ []byte, _ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (k *G1Elt) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *G1Elt) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (k *G1Elt) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (k *G1Elt) Neg(a kyber.Point) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (k *G1Elt) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// MarshalBinary returns a compressed point, without any domain separation tag information
func (k *G1Elt) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	// we need to clone the point because of https://github.com/kilic/bls12-381/issues/37
	// in order to avoid risks of race conditions.
	return nil, nil
}

// UnmarshalBinary populates the point from a compressed point representation.
func (k *G1Elt) UnmarshalBinary(buff []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalTo writes a compressed point to the Writer, without any domain separation tag information
func (k *G1Elt) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// UnmarshalFrom populates the point from a compressed point representation read from the Reader.
func (k *G1Elt) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (k *G1Elt) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (k *G1Elt) String() string { _ = "STUB: not implemented"; return "" }

func (k *G1Elt) Hash(m []byte) kyber.Point {
	_ = "STUB: not implemented"

	// We treat a 0 len dst as the default value as per the RFC "Tags MUST have nonzero length"
	return *new(kyber.Point)
}

func (k *G1Elt) IsInCorrectGroup() bool { _ = "STUB: not implemented"; return false }
