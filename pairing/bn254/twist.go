//go:build !constantTime

package bn254

import (
	"math/big"
)

// twistPoint implements the elliptic curve y²=x³+3/ξ over GF(p²). Points are
// kept in Jacobian form and t=z² when valid. The group G₂ is the set of
// n-torsion points of this curve over GF(p²) (where n = Order)
type twistPoint struct {
	x, y, z, t gfP2
}

var twistB = &gfP2{
	gfP{0x38e7ecccd1dcff67, 0x65f0b37d93ce0d3e, 0xd749d0dd22ac00aa, 0x0141b9ce4a688d4d},
	gfP{0x3bf938e377b802a8, 0x020b1b273633535d, 0x26b7edf049755260, 0x2514c6324384a86d},
}

// twistGen is the generator of group G₂.
var twistGen = &twistPoint{
	gfP2{
		gfP{0xafb4737da84c6140, 0x6043dd5a5802d8c4, 0x09e950fc52a02f86, 0x14fef0833aea7b6b},
		gfP{0x8e83b5d102bc2026, 0xdceb1935497b0172, 0xfbb8264797811adf, 0x19573841af96503b},
	},
	gfP2{
		gfP{0x64095b56c71856ee, 0xdc57f922327d3cbb, 0x55f935be33351076, 0x0da4a0e693fd6482},
		gfP{0x619dfa9d886be9f6, 0xfe7fd297f59e9b78, 0xff9e1a62231b7dfe, 0x28fd7eebae9e4206},
	},
	gfP2{*newGFp(0), *newGFp(1)},
	gfP2{*newGFp(0), *newGFp(1)},
}

func (c *twistPoint) String() string { _ = "STUB: not implemented"; return "" }

func (c *twistPoint) Set(a *twistPoint) { _ = "STUB: not implemented"; return }

// IsOnCurve returns true iff c is on the curve.
func (c *twistPoint) IsOnCurve() bool { _ = "STUB: not implemented"; return false }

func (c *twistPoint) SetInfinity() { _ = "STUB: not implemented"; return }

func (c *twistPoint) IsInfinity() bool { _ = "STUB: not implemented"; return false }

func (c *twistPoint) Add(a, b *twistPoint) {
	_ = "STUB: not implemented"
	// For additional comments, see the same function in curve.go.
	return
}

// See http://hyperelliptic.org/EFD/g1p/auto-code/shortw/jacobian-0/addition/add-2007-bl.op3

// t7
// t8
// t9
// t10

// t11
// t12
// t13
// t14

func (c *twistPoint) Double(a *twistPoint) {
	_ = "STUB: not implemented"
	// See http://hyperelliptic.org/EFD/g1p/auto-code/shortw/jacobian-0/doubling/dbl-2009-l.op3
	return
}

func (c *twistPoint) Mul(a *twistPoint, scalar *big.Int) { _ = "STUB: not implemented"; return }

// NB: Not safe for concurrent calls
func (c *twistPoint) MakeAffine() { _ = "STUB: not implemented"; return }

func (c *twistPoint) Neg(a *twistPoint) { _ = "STUB: not implemented"; return }

// Clone makes a deep copy of the point
func (c *twistPoint) Clone() *twistPoint { _ = "STUB: not implemented"; return nil }
