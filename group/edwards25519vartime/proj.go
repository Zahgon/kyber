//go:build !constantTime

package edwards25519vartime

import (
	"crypto/cipher"
	"errors"

	"io"

	"go.dedis.ch/kyber/v4/compatible"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/group/mod"
)

var ErrTypeCast = errors.New("invalid type cast")

type projPoint struct {
	X, Y, Z mod.Int
	c       *ProjectiveCurve
}

func (P *projPoint) initXY(x, y *compatible.Int, c kyber.Group) { _ = "STUB: not implemented"; return }

//nolint:errcheck // Design pattern to emulate generics

func (P *projPoint) getXY() (x, y *mod.Int) { _ = "STUB: not implemented"; return nil, nil }

func (P *projPoint) String() string { _ = "STUB: not implemented"; return "" }

func (P *projPoint) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (P *projPoint) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (P *projPoint) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (P *projPoint) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (P *projPoint) UnmarshalFrom(r io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Equality test for two Points on the same curve.
// We can avoid inversions here because:
//
//	(X1/Z1,Y1/Z1) == (X2/Z2,Y2/Z2)
//		iff
//	(X1*Z2,Y1*Z2) == (X2*Z1,Y2*Z1)
func (P *projPoint) Equal(CP2 kyber.Point) bool { _ = "STUB: not implemented"; return false }

//nolint:errcheck // Design pattern to emulate generics

func (P *projPoint) Set(CP2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *
	//nolint:errcheck // Design pattern to emulate generics
	new(kyber.Point)
}

func (P *projPoint) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *projPoint) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *projPoint) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *projPoint) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

// Normalize the point's representation to Z=1.
func (P *projPoint) normalize() { _ = "STUB: not implemented"; return }

func (P *projPoint) Embed(data []byte, rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (P *projPoint) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Extract embedded data from a point group element
func (P *projPoint) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Add two points using optimized projective coordinate addition formulas.
// Formulas taken from:
//
//	http://eprint.iacr.org/2008/013.pdf
//	https://hyperelliptic.org/EFD/g1p/auto-twisted-projective.html
//
//nolint:dupl //Doesn't make sense to extract part of Add(), Sub()
func (P *projPoint) Add(CP1, CP2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *
	//nolint:errcheck // Design pattern to emulate generics
	new(kyber.Point)
}

//nolint:errcheck // Design pattern to emulate generics

// Subtract points so that their scalars subtract homomorphically
//
//nolint:dupl //Doesn't make sense to extract part of Add(), Sub(), double()
func (P *projPoint) Sub(CP1, CP2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *
	//nolint:errcheck // Design pattern to emulate generics
	new(kyber.Point)
}

//nolint:errcheck // Design pattern to emulate generics

// Find the negative of point A.
// For Edwards curves, the negative of (x,y) is (-x,y).
func (P *projPoint) Neg(CA kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Point)
}

// Optimized point doubling for use in scalar multiplication.
func (P *projPoint) double() { _ = "STUB: not implemented"; return }

// Multiply point p by scalar s using the repeated doubling method.
func (P *projPoint) Mul(s kyber.Scalar, G kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Must use temporary for in-place multiply

// Initialize to identity element (0,1)

// ProjectiveCurve implements Twisted Edwards curves
// using projective coordinate representation (X:Y:Z),
// satisfying the identities x = X/Z, y = Y/Z.
// This representation still supports all Twisted Edwards curves
// and avoids expensive modular inversions on the critical paths.
// Uses the projective arithmetic formulas in:
// http://cr.yp.to/newelliptic/newelliptic-20070906.pdf
type ProjectiveCurve struct {
	curve           // generic Edwards curve functionality
	null  projPoint // Constant identity/null point (0,1)
	base  projPoint // Standard base point
}

// Point creates a new Point on this curve.
func (c *ProjectiveCurve) Point() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

// Init initializes the curve with given parameters.
func (c *ProjectiveCurve) Init(p *Param, fullGroup bool) *ProjectiveCurve {
	_ = "STUB: not implemented"
	return nil
}
