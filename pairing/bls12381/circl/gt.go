package circl

import (
	"crypto/cipher"
	"io"

	bls12381 "github.com/cloudflare/circl/ecc/bls12381"
	"go.dedis.ch/kyber/v4"
)

var gtBase *bls12381.Gt

func init() {
	gtBase = bls12381.Pair(bls12381.G1Generator(), bls12381.G2Generator())
}

var _ kyber.Point = &GTElt{}

// GTElt is a wrapper around the Circl Gt point type.
type GTElt struct{ inner bls12381.Gt }

// MarshalBinary returns a compressed point, without any domain separation tag information
func (p *GTElt) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil

	// UnmarshalBinary populates the point from a compressed point representation.
}

func (p *GTElt) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (p *GTElt) String() string { _ = "STUB: not implemented"; return "" }

func (p *GTElt) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

// MarshalTo writes a compressed point to the Writer, without any domain separation tag information
func (p *GTElt) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// UnmarshalFrom populates the point from a compressed point representation read from the Reader.
func (p *GTElt) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *GTElt) Equal(p2 kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (p *GTElt) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *GTElt) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *GTElt) Pick(_ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *GTElt) Set(p2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *GTElt) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *GTElt) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (p *GTElt) Embed(_ []byte, _ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *GTElt) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *GTElt) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *GTElt) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (p *GTElt) Neg(a kyber.Point) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (p *GTElt) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}
