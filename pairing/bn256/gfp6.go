//go:build !constantTime

package bn256

// For details of the algorithms used, see "Multiplication and Squaring on
// Pairing-Friendly Fields, Devegili et al.
// http://eprint.iacr.org/2006/471.pdf.

// gfP6 implements the field of size p⁶ as a cubic extension of gfP2 where τ³=ξ
// and ξ=i+3.
type gfP6 struct {
	x, y, z gfP2 // value is xτ² + yτ + z
}

func (e *gfP6) String() string { _ = "STUB: not implemented"; return "" }

func (e *gfP6) Set(a *gfP6) *gfP6 { _ = "STUB: not implemented"; return nil }

func (e *gfP6) SetZero() *gfP6 { _ = "STUB: not implemented"; return nil }

func (e *gfP6) SetOne() *gfP6 { _ = "STUB: not implemented"; return nil }

func (e *gfP6) IsZero() bool { _ = "STUB: not implemented"; return false }

func (e *gfP6) IsOne() bool { _ = "STUB: not implemented"; return false }

func (e *gfP6) Neg(a *gfP6) *gfP6 { _ = "STUB: not implemented"; return nil }

func (e *gfP6) Frobenius(a *gfP6) *gfP6 { _ = "STUB: not implemented"; return nil }

// FrobeniusP2 computes (xτ²+yτ+z)^(p²) = xτ^(2p²) + yτ^(p²) + z
func (e *gfP6) FrobeniusP2(a *gfP6) *gfP6 {
	_ = "STUB: not implemented"
	// τ^(2p²) = τ²τ^(2p²-2) = τ²ξ^((2p²-2)/3)
	return nil
}

// τ^(p²) = ττ^(p²-1) = τξ^((p²-1)/3)

func (e *gfP6) FrobeniusP4(a *gfP6) *gfP6 { _ = "STUB: not implemented"; return nil }

func (e *gfP6) Add(a, b *gfP6) *gfP6 { _ = "STUB: not implemented"; return nil }

func (e *gfP6) Sub(a, b *gfP6) *gfP6 { _ = "STUB: not implemented"; return nil }

func (e *gfP6) Mul(a, b *gfP6) *gfP6 {
	_ = "STUB: not implemented"
	// "Multiplication and Squaring on Pairing-Friendly Fields"
	// Section 4, Karatsuba method.
	// http://eprint.iacr.org/2006/471.pdf
	return nil
}

func (e *gfP6) MulScalar(a *gfP6, b *gfP2) *gfP6 { _ = "STUB: not implemented"; return nil }

func (e *gfP6) MulGFP(a *gfP6, b *gfP) *gfP6 { _ = "STUB: not implemented"; return nil }

// MulTau computes τ·(aτ²+bτ+c) = bτ²+cτ+aξ
func (e *gfP6) MulTau(a *gfP6) *gfP6 { _ = "STUB: not implemented"; return nil }

func (e *gfP6) Square(a *gfP6) *gfP6 { _ = "STUB: not implemented"; return nil }

func (e *gfP6) Invert(a *gfP6) *gfP6 {
	_ = "STUB: not implemented"
	// See "Implementing cryptographic pairings", M. Scott, section 3.2.
	// ftp://136.206.11.249/pub/crypto/pairings.pdf
	return nil
}

// Here we can give a short explanation of how it works: let j be a cubic root of
// unity in GF(p²) so that 1+j+j²=0.
// Then (xτ² + yτ + z)(xj²τ² + yjτ + z)(xjτ² + yj²τ + z)
// = (xτ² + yτ + z)(Cτ²+Bτ+A)
// = (x³ξ²+y³ξ+z³-3ξxyz) = F is an element of the base field (the norm).
//
// On the other hand (xj²τ² + yjτ + z)(xjτ² + yj²τ + z)
// = τ²(y²-ξxz) + τ(ξx²-yz) + (z²-ξxy)
//
// So that's why A = (z²-ξxy), B = (ξx²-yz), C = (y²-ξxz)

// Clone makes a hard copy of the field
func (e *gfP6) Clone() gfP6 { _ = "STUB: not implemented"; return *new(gfP6) }
