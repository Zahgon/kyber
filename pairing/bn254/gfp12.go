//go:build !constantTime

package bn254

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

func (e *gfP12) MulScalar(a *gfP6) *gfP12 { _ = "STUB: not implemented"; return nil }

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

// Clone makes a deep copy of the field
func (e *gfP12) Clone() *gfP12 { _ = "STUB: not implemented"; return nil }
