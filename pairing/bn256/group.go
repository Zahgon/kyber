//go:build !constantTime

package bn256

import (
	"crypto/cipher"

	"go.dedis.ch/kyber/v4"
)

type groupG1 struct {
	common
	*commonSuite
}

func (g *groupG1) String() string { _ = "STUB: not implemented"; return "" }

func (g *groupG1) PointLen() int { _ = "STUB: not implemented"; return 0 }

func (g *groupG1) Point() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

type groupG2 struct {
	common
	*commonSuite
}

func (g *groupG2) String() string { _ = "STUB: not implemented"; return "" }

func (g *groupG2) PointLen() int { _ = "STUB: not implemented"; return 0 }

func (g *groupG2) Point() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

type groupGT struct {
	common
	*commonSuite
}

func (g *groupGT) String() string { _ = "STUB: not implemented"; return "" }

func (g *groupGT) PointLen() int { _ = "STUB: not implemented"; return 0 }

func (g *groupGT) Point() kyber.Point {
	_ = "STUB: not implemented"
	return *

	// common functionalities across G1, G2, and GT
	new(kyber.Point)
}

type common struct{}

func (c *common) ScalarLen() int { _ = "STUB: not implemented"; return 0 }

func (c *common) Scalar() kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }

func (c *common) PrimeOrder() bool { _ = "STUB: not implemented"; return false }

func (c *common) NewKey(rand cipher.Stream) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}
