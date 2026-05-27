//go:build !constantTime

package edwards25519vartime

import (
	"crypto/cipher"
	"io"

	"go.dedis.ch/kyber/v4/compatible"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/group/mod"
)

type extPoint struct {
	X, Y, Z, T mod.Int
	c          *ExtendedCurve
}

func (P *extPoint) initXY(x, y *compatible.Int, c kyber.Group) { _ = "STUB: not implemented"; return }

//nolint:errcheck // Design pattern to emulate generics

func (P *extPoint) getXY() (x, y *mod.Int) { _ = "STUB: not implemented"; return nil, nil }

func (P *extPoint) String() string { _ = "STUB: not implemented"; return "" }

func (P *extPoint) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (P *extPoint) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (P *extPoint) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (P *extPoint) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (P *extPoint) UnmarshalFrom(r io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Equality test for two Points on the same curve.
// We can avoid inversions here because:
//
//	(X1/Z1,Y1/Z1) == (X2/Z2,Y2/Z2)
//		iff
//	(X1*Z2,Y1*Z2) == (X2*Z1,Y2*Z1)
func (P *extPoint) Equal(CP2 kyber.Point) bool {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return false
}

func (P *extPoint) Set(CP2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *
	//nolint:errcheck // Design pattern to emulate generics
	new(kyber.Point)
}

func (P *extPoint) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *extPoint) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *extPoint) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *extPoint) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

// Normalize the point's representation to Z=1.
func (P *extPoint) normalize() { _ = "STUB: not implemented"; return }

// Check the validity of the T coordinate
//
//nolint:unused // may be useful
func (P *extPoint) checkT() { _ = "STUB: not implemented"; return }

func (P *extPoint) Embed(data []byte, rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (P *extPoint) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Extract embedded data from a point group element
func (P *extPoint) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Add two points using optimized extended coordinate addition formulas.
//
//nolint:dupl //Doesn't make sense to extract part of Add(), Sub(), double()
func (P *extPoint) Add(CP1, CP2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *
	//nolint:errcheck // Design pattern to emulate generics
	new(kyber.Point)
}

//nolint:errcheck // Design pattern to emulate generics

// Subtract points.
//
//nolint:dupl //Doesn't make sense to extract part of Add(), Sub(), double()
func (P *extPoint) Sub(CP1, CP2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *
	//nolint:errcheck // Design pattern to emulate generics
	new(kyber.Point)
}

//nolint:errcheck // Design pattern to emulate generics

// Find the negative of point A.
// For Edwards curves, the negative of (x,y) is (-x,y).
func (P *extPoint) Neg(CA kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Point)
}

// Optimized point doubling for use in scalar multiplication.
// Uses the formulae in section 3.3 of:
// https://www.iacr.org/archive/asiacrypt2008/53500329/53500329.pdf
func (P *extPoint) double() { _ = "STUB: not implemented"; return }

// Multiply point p by scalar s using the repeated doubling method.
//
// Currently doesn't implement the optimization of
// switching between projective and extended coordinates during
// scalar multiplication.
func (P *extPoint) Mul(s kyber.Scalar, G kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Must use temporary for in-place multiply

// Initialize to identity element (0,1)

// ExtendedCurve implements Twisted Edwards curves
// using projective coordinate representation (X:Y:Z),
// satisfying the identities x = X/Z, y = Y/Z.
// This representation still supports all Twisted Edwards curves
// and avoids expensive modular inversions on the critical paths.
// Uses the projective arithmetic formulas in:
// http://cr.yp.to/newelliptic/newelliptic-20070906.pdf
//

// ExtendedCurve implements Twisted Edwards curves
// using the Extended Coordinate representation specified in:
// Hisil et al, "Twisted Edwards Curves Revisited",
// http://eprint.iacr.org/2008/522
//
// This implementation is designed to work with all Twisted Edwards curves,
// foregoing the further optimizations that are available for the
// special case with curve parameter a=-1.
// We leave the task of hyperoptimization to curve-specific implementations
// such as the ed25519 package.
type ExtendedCurve struct {
	curve          // generic Edwards curve functionality
	null  extPoint // Constant identity/null point (0,1)
	base  extPoint // Standard base point
}

// Point creates a new Point on this curve.
func (c *ExtendedCurve) Point() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

// InitCurve initializes the curve with given parameters.
func (c *ExtendedCurve) InitCurve(p *Param, fullGroup bool) *ExtendedCurve {
	_ = "STUB: not implemented"
	return nil
}
