//go:build !constantTime

package bn256

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
	gfP{0x75046774386b8d71, 0x5bd0854a46d36cf8, 0x664327a1d41c8414, 0x96c9abb932eeb2f},
	gfP{0xb94f760fb4c5ee14, 0xdae9f8f24c3b6eb4, 0x77a675d2e52f4fe4, 0x736f31b09116c66b},
}

// twistGen is the generator of group G₂.
var twistGen = &twistPoint{
	gfP2{
		gfP{0x402c4ab7139e1404, 0xce1c368a183d85a4, 0xd67cf9a6cb8d3983, 0x3cf246bbc2a9fbe8},
		gfP{0x88f9f11da7cdc184, 0x18293f95d69509d3, 0xb5ce0c55a735d5a1, 0x15134189bfd45a0},
	},
	gfP2{
		gfP{0xbfac7d731e9e87a2, 0xa50bb8007962e441, 0xafe910a4e8270556, 0x5075c5429d69159a},
		gfP{0xc2e07c1463ea9e56, 0xee4442052072ebd2, 0x561a519486036937, 0x5bd9394cc0d2cce},
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

func (c *twistPoint) MakeAffine() { _ = "STUB: not implemented"; return }

func (c *twistPoint) Neg(a *twistPoint) { _ = "STUB: not implemented"; return }

// Clone makes a hard copy of the point
func (c *twistPoint) Clone() *twistPoint { _ = "STUB: not implemented"; return nil }
