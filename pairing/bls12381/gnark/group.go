//go:build !constantTime

package gnark

import (
	"go.dedis.ch/kyber/v4"
)

var (
	G1 kyber.Group = &groupBls{name: "bls12-381.G1", newPoint: func() kyber.Point { return new(G1Elt).Null() }}
	G2 kyber.Group = &groupBls{name: "bls12-381.G2", newPoint: func() kyber.Point { return new(G2Elt).Null() }}
	GT kyber.Group = &groupBls{name: "bls12-381.GT", newPoint: func() kyber.Point { return new(GTElt).Null() }}
)

type groupBls struct {
	name     string
	newPoint func() kyber.Point
}

func (g groupBls) String() string       { _ = "STUB: not implemented"; return "" }
func (g groupBls) ScalarLen() int       { _ = "STUB: not implemented"; return 0 }
func (g groupBls) Scalar() kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }
func (g groupBls) PointLen() int        { _ = "STUB: not implemented"; return 0 }
func (g groupBls) Point() kyber.Point   { _ = "STUB: not implemented"; return *new(kyber.Point) }
