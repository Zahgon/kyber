//go:build !constantTime

package kilic

import (
	"crypto/cipher"
	"io"

	bls12381 "github.com/kilic/bls12-381"
	"go.dedis.ch/kyber/v4"
)

// GTElt contains a Gt element from the Kilic BLS12-381 curve
type GTElt struct {
	f *bls12381.E
}

func newEmptyGT() *GTElt { _ = "STUB: not implemented"; return nil }

func newGT(f *bls12381.E) *GTElt { _ = "STUB: not implemented"; return nil }

func (k *GTElt) Equal(kk kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (k *GTElt) Null() kyber.Point {
	_ = "STUB: not implemented"
	// One since we deal with Gt elements as a multiplicative group only
	// i.e. Add in kyber -> mul in kilic/, Neg in kyber -> inverse in kilic/ etc
	return *new(kyber.Point)
}

func (k *GTElt) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (k *GTElt) Pick(_ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (k *GTElt) Set(q kyber.Point) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (k *GTElt) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (k *GTElt) Add(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (k *GTElt) Sub(a, b kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (k *GTElt) Neg(q kyber.Point) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (k *GTElt) Mul(s kyber.Scalar, q kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// MarshalBinary returns a compressed point, without any domain separation tag information
func (k *GTElt) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalTo writes a compressed point to the Writer, without any domain separation tag information
func (k *GTElt) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// UnmarshalBinary populates the point from a compressed point representation.
func (k *GTElt) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalFrom populates the point from a compressed point representation read from the Reader.
func (k *GTElt) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (k *GTElt) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (k *GTElt) String() string { _ = "STUB: not implemented"; return "" }

func (k *GTElt) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (k *GTElt) Embed(_ []byte, _ cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (k *GTElt) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
