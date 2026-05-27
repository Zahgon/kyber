//go:build !constantTime

package kilic

import (
	"crypto/cipher"
	"hash"

	"go.dedis.ch/kyber/v4"
)

type groupBls struct {
	str      string
	newPoint func() kyber.Point
	isPrime  bool
}

func (g *groupBls) String() string { _ = "STUB: not implemented"; return "" }

func (g *groupBls) Scalar() kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }

func (g *groupBls) ScalarLen() int { _ = "STUB: not implemented"; return 0 }

func (g *groupBls) PointLen() int { _ = "STUB: not implemented"; return 0 }

func (g *groupBls) Point() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (g *groupBls) IsPrimeOrder() bool { _ = "STUB: not implemented"; return false }

func (g *groupBls) Hash() hash.Hash {
	_ = "STUB: not implemented"
	return *

	// XOF returns a newly instantiated blake2xb XOF function.
	new(hash.Hash)
}

func (g *groupBls) XOF(seed []byte) kyber.XOF {
	_ = "STUB: not implemented"
	return *

	// RandomStream returns a cipher.Stream which corresponds to a key stream from
	// crypto/rand.
	new(kyber.XOF)
}

func (g *groupBls) RandomStream() cipher.Stream {
	_ = "STUB: not implemented"
	return *new(cipher.Stream)
}

func NewGroupG1(dst ...byte) kyber.Group { _ = "STUB: not implemented"; return *new(kyber.Group) }

func NewGroupG2(dst ...byte) kyber.Group { _ = "STUB: not implemented"; return *new(kyber.Group) }

func NewGroupGT() kyber.Group { _ = "STUB: not implemented"; return *new(kyber.Group) }
