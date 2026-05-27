//go:build experimental && !constantTime

package edwards25519vartime

import (
	"crypto/cipher"
	"io"

	"io"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/compatible"
	"go.dedis.ch/kyber/v4/group/mod"
)

type basicPoint struct {
	x, y mod.Int
	c    *BasicCurve
}

func (P *basicPoint) initXY(x, y *compatible.Int, c kyber.Group) { _ = "STUB: not implemented"; return }

func (P *basicPoint) getXY() (x, y *mod.Int) { _ = "STUB: not implemented"; return nil, nil }

func (P *basicPoint) String() string { _ = "STUB: not implemented"; return "" }

// coord creates a new ModInt representing a coordinate on this curve,
// with a given int64 integer value for constant-initialization convenience.
func (P *basicPoint) coord(v int64) *mod.Int { _ = "STUB: not implemented"; return nil }

func (P *basicPoint) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

// MarshalBinary encodew an Edwards curve point.
func (P *basicPoint) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary decodes an Edwards curve point.
func (P *basicPoint) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (P *basicPoint) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (P *basicPoint) UnmarshalFrom(r io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Equal tests for two Points on the same curve
func (P *basicPoint) Equal(P2 kyber.Point) bool { _ = "STUB: not implemented"; return false }

// Set point to be equal to P2.
func (P *basicPoint) Set(P2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Clone returns the given point
func (P *basicPoint) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

// Null sets to the neutral element, which is (0,1) for twisted Edwards curves.
func (P *basicPoint) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

// Base sets to the standard base point for this curve
func (P *basicPoint) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *basicPoint) EmbedLen() int { _ = "STUB: not implemented"; return 0 }

func (P *basicPoint) Embed(data []byte, rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (P *basicPoint) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Data extracts embedded data from a point group element
func (P *basicPoint) Data() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Add two points using the basic unified addition laws for Edwards curves:
		//
		//	x' = ((x1*y2 + x2*y1) / (1 + d*x1*x2*y1*y2))
		//	y' = ((y1*y2 - a*x1*x2) / (1 - d*x1*x2*y1*y2))
		nil
}

func (P *basicPoint) Add(P1, P2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Reused part of denominator: dm = d*x1*x2*y1*y2

// x' numerator/denominator

// y' numerator/denominator

// result point

// Point doubling, which for Edwards curves can be accomplished
// simply by adding a point to itself (no exceptions for equal input points).
func (P *basicPoint) double() kyber.Point {
	_ = "STUB: not implemented"

	// Subtract points so that their scalars subtract homomorphically
	return *new(kyber.Point)
}

func (P *basicPoint) Sub(A, B kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Find the negative of point A.
// For Edwards curves, the negative of (x,y) is (-x,y).
func (P *basicPoint) Neg(A kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Multiply point p by scalar s using the repeated doubling method.
func (P *basicPoint) Mul(s kyber.Scalar, G kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Must use temporary in case G == P

// Initialize to identity element (0,1)

// Basic unoptimized reference implementation of Twisted Edwards curves.
// This reference implementation is mainly intended for testing, debugging,
// and instructional uses, and not for production use.
// The projective coordinates implementation (ProjectiveCurve)
// is just as general and much faster.
type BasicCurve struct {
	curve            // generic Edwards curve functionality
	null  basicPoint // Neutral/identity point (0,1)
	base  basicPoint // Standard base point
}

// Create a new Point on this curve.
func (c *BasicCurve) Point() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

// Initialize the curve with given parameters.
func (c *BasicCurve) Init(p *Param, fullGroup bool) *BasicCurve {
	_ = "STUB: not implemented"
	return nil
}
