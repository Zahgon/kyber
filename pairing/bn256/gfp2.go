//go:build !constantTime

package bn256

// For details of the algorithms used, see "Multiplication and Squaring on
// Pairing-Friendly Fields, Devegili et al.
// http://eprint.iacr.org/2006/471.pdf.

// gfP2 implements a field of size p² as a quadratic extension of the base field
// where i²=-1.
type gfP2 struct {
	x, y gfP // value is xi+y.
}

func gfP2Decode(in *gfP2) *gfP2 { _ = "STUB: not implemented"; return nil }

func (e *gfP2) String() string { _ = "STUB: not implemented"; return "" }

func (e *gfP2) Set(a *gfP2) *gfP2 { _ = "STUB: not implemented"; return nil }

func (e *gfP2) SetZero() *gfP2 { _ = "STUB: not implemented"; return nil }

func (e *gfP2) SetOne() *gfP2 { _ = "STUB: not implemented"; return nil }

func (e *gfP2) IsZero() bool { _ = "STUB: not implemented"; return false }

func (e *gfP2) IsOne() bool { _ = "STUB: not implemented"; return false }

func (e *gfP2) Conjugate(a *gfP2) *gfP2 { _ = "STUB: not implemented"; return nil }

func (e *gfP2) Neg(a *gfP2) *gfP2 { _ = "STUB: not implemented"; return nil }

func (e *gfP2) Add(a, b *gfP2) *gfP2 { _ = "STUB: not implemented"; return nil }

func (e *gfP2) Sub(a, b *gfP2) *gfP2 { _ = "STUB: not implemented"; return nil }

// See "Multiplication and Squaring in Pairing-Friendly Fields",
// http://eprint.iacr.org/2006/471.pdf
func (e *gfP2) Mul(a, b *gfP2) *gfP2 { _ = "STUB: not implemented"; return nil }

func (e *gfP2) MulScalar(a *gfP2, b *gfP) *gfP2 { _ = "STUB: not implemented"; return nil }

// MulXi sets e=ξa where ξ=i+3 and then returns e.
func (e *gfP2) MulXi(a *gfP2) *gfP2 {
	_ = "STUB: not implemented"
	// (xi+y)(i+3) = (3x+y)i+(3y-x)
	return nil
}

func (e *gfP2) Square(a *gfP2) *gfP2 {
	_ = "STUB: not implemented"
	// Complex squaring algorithm:
	// (xi+y)² = (x+y)(y-x) + 2*i*x*y
	return nil
}

func (e *gfP2) Invert(a *gfP2) *gfP2 {
	_ = "STUB: not implemented"
	// See "Implementing cryptographic pairings", M. Scott, section 3.2.
	// ftp://136.206.11.249/pub/crypto/pairings.pdf
	return nil
}

// Clone makes a hard copy of the field
func (e *gfP2) Clone() gfP2 { _ = "STUB: not implemented"; return *new(gfP2) }
