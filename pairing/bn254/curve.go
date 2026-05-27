//go:build !constantTime

package bn254

import (
	"math/big"
)

// curvePoint implements the elliptic curve y²=x³+3. Points are kept in Jacobian
// form and t=z² when valid. G₁ is the set of points of this curve on GF(p).
type curvePoint struct {
	x, y, z, t gfP
}

// curveGen is the generator of G₁.
var curveGen = &curvePoint{
	x: *newGFp(1),
	y: *newGFp(2),
	z: *newGFp(1),
	t: *newGFp(1),
}

// evaluate the curve at x
func g(x *gfP) *gfP { _ = "STUB: not implemented"; return nil }

func (c *curvePoint) String() string { _ = "STUB: not implemented"; return "" }

func (c *curvePoint) Set(a *curvePoint) { _ = "STUB: not implemented"; return }

// IsOnCurve returns true iff c is on the curve.
func (c *curvePoint) IsOnCurve() bool { _ = "STUB: not implemented"; return false }

func (c *curvePoint) SetInfinity() { _ = "STUB: not implemented"; return }

func (c *curvePoint) IsInfinity() bool { _ = "STUB: not implemented"; return false }

func (c *curvePoint) Add(a, b *curvePoint) { _ = "STUB: not implemented"; return }

// See http://hyperelliptic.org/EFD/g1p/auto-code/shortw/jacobian-0/addition/add-2007-bl.op3

// Normalize the points by replacing a = [x1:y1:z1] and b = [x2:y2:z2]
// by [u1:s1:z1·z2] and [u2:s2:z1·z2]
// where u1 = x1·z2², s1 = y1·z2³ and u1 = x2·z1², s2 = y2·z1³

// Compute x = (2h)²(s²-u1-u2)
// where s = (s2-s1)/(u2-u1) is the slope of the line through
// (u1,s1) and (u2,s2). The extra factor 2h = 2(u2-u1) comes from the value of z below.
// This is also:
// 4(s2-s1)² - 4h²(u1+u2) = 4(s2-s1)² - 4h³ - 4h²(2u1)
//                        = r² - j - 2v
// with the notations below.

// i = 4h²

// j = 4h³

// t4 = 4(s2-s1)²

// Set y = -(2h)³(s1 + s*(x/4h²-u1))
// This is also
// y = - 2·s1·j - (s2-s1)(2x - 2i·u1) = r(v-x) - 2·s1·j
// t7
// t8
// t9
// t10

// Set z = 2(u2-u1)·z1·z2 = 2h·z1·z2
// t11
// t12
// t13
// t14

func (c *curvePoint) Double(a *curvePoint) {
	_ = "STUB: not implemented"
	// See http://hyperelliptic.org/EFD/g1p/auto-code/shortw/jacobian-0/doubling/dbl-2009-l.op3
	return
}

func (c *curvePoint) Mul(a *curvePoint, scalar *big.Int) { _ = "STUB: not implemented"; return }

func (c *curvePoint) MakeAffine() { _ = "STUB: not implemented"; return }

func (c *curvePoint) Neg(a *curvePoint) { _ = "STUB: not implemented"; return }

// Clone makes a deep copy of the curve point
func (c *curvePoint) Clone() *curvePoint { _ = "STUB: not implemented"; return nil }
