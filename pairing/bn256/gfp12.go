//go:build !constantTime

package bn256

// For details of the algorithms used, see "Multiplication and Squaring on
// Pairing-Friendly Fields, Devegili et al.
// http://eprint.iacr.org/2006/471.pdf.

import (
	"math/big"
)

// gfP12 implements the field of size p¹² as a quadratic extension of gfP6
// where ω²=τ.
type gfP12 struct {
	x, y gfP6 // value is xω + y
}

var gfP12Gen = &gfP12{
	x: gfP6{
		x: gfP2{
			x: gfP{0x62d608d6bb67a4fb, 0x9a66ec93f0c2032f, 0x5391628e924e1a34, 0x2162dbf7de801d0e},
			y: gfP{0x3e0c1a72bf08eb4f, 0x4972ec05990a5ecc, 0xf7b9a407ead8007e, 0x3ca04c613572ce49},
		},
		y: gfP2{
			x: gfP{0xace536a5607c910e, 0xda93774a941ddd40, 0x5de0e9853b7593ad, 0xe05bb926f513153},
			y: gfP{0x3f4c99f8abaf1a22, 0x66d5f6121f86dc33, 0x8e0a82f68a50abba, 0x819927d1eebd0695},
		},
		z: gfP2{
			x: gfP{0x7cdef49c5477faa, 0x40eb71ffedaa199d, 0xbc896661f17c9b8f, 0x3144462983c38c02},
			y: gfP{0xcd09ee8dd8418013, 0xf8d050d05faa9b11, 0x589e90a555507ee1, 0x58e4ab25f9c49c15},
		},
	},
	y: gfP6{
		x: gfP2{
			x: gfP{0x7e76809b142d020b, 0xd9949d1b2822e995, 0x3de93d974f84b076, 0x144523477028928d},
			y: gfP{0x79952799f9ef4b0, 0x4102c47aa3df01c6, 0xfa82a633c53da2e1, 0x54c3f0392f9f7e0e},
		},
		y: gfP2{
			x: gfP{0xd3432a335533272b, 0xa008fbbdc7d74f4a, 0x68e3c81eb7295ed9, 0x17fe34c21fdecef2},
			y: gfP{0xfb0bc4c0ef6df55f, 0x8bdc585b70bc2120, 0x17d498d2cb720def, 0x2a368248319b899c},
		},
		z: gfP2{
			x: gfP{0xf8487d81cb354c6c, 0x7421be69f1522caa, 0x6940c778b9fb2d54, 0x7da4b04e102bb621},
			y: gfP{0x97b91989993e7be4, 0x8526545356eab684, 0xb050073022eb1892, 0x658b432ad09939c0},
		},
	},
}

var gfP12Inf = &gfP12{
	x: gfP6{
		x: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		},
		y: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		},
		z: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		},
	},
	y: gfP6{
		x: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		},
		y: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		},
		z: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0xe7a35393a1f76999, 0x11a4772edf4a4a61, 0x559013479e7b23de, 0x704afe1cb55c7806},
		},
	},
}

func (e *gfP12) String() string { _ = "STUB: not implemented"; return "" }

func (e *gfP12) Set(a *gfP12) *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) SetZero() *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) SetOne() *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) IsZero() bool { _ = "STUB: not implemented"; return false }

func (e *gfP12) IsOne() bool { _ = "STUB: not implemented"; return false }

func (e *gfP12) Conjugate(a *gfP12) *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) Neg(a *gfP12) *gfP12 { _ = "STUB: not implemented"; return nil }

// Frobenius computes (xω+y)^p = x^p ω·ξ^((p-1)/6) + y^p
func (e *gfP12) Frobenius(a *gfP12) *gfP12 { _ = "STUB: not implemented"; return nil }

// FrobeniusP2 computes (xω+y)^p² = x^p² ω·ξ^((p²-1)/6) + y^p²
func (e *gfP12) FrobeniusP2(a *gfP12) *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) FrobeniusP4(a *gfP12) *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) Add(a, b *gfP12) *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) Sub(a, b *gfP12) *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) Mul(a, b *gfP12) *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) MulScalar(b *gfP6) *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) Exp(a *gfP12, power *big.Int) *gfP12 { _ = "STUB: not implemented"; return nil }

func (e *gfP12) Square(a *gfP12) *gfP12 {
	_ = "STUB: not implemented"
	// Complex squaring algorithm
	return nil
}

func (e *gfP12) Invert(a *gfP12) *gfP12 {
	_ = "STUB: not implemented"
	// See "Implementing cryptographic pairings", M. Scott, section 3.2.
	// ftp://136.206.11.249/pub/crypto/pairings.pdf
	return nil
}

// Clone makes a hard copy of the field
func (e *gfP12) Clone() *gfP12 { _ = "STUB: not implemented"; return nil }
