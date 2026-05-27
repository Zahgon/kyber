//go:build !constantTime

package p256

import (
	"crypto/cipher"
	"crypto/elliptic"
	"io"
	"math/big"

	"go.dedis.ch/kyber/v4"
)

type curvePoint struct {
	x, y *big.Int
	c    *curve
}

func (P *curvePoint) String() string { _ = "STUB: not implemented"; return "" }

func (P *curvePoint) Equal(P2 kyber.Point) bool { _ = "STUB: not implemented"; return false }

//nolint:errcheck // Design pattern to emulate generics

// Make sure both coordinates are normalized.
// Apparently Go's elliptic curve code doesn't always ensure this.
// Use temporary big.Ints to avoid mutating the operands.

func (P *curvePoint) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *curvePoint) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *curvePoint) Valid() bool {
	_ = "STUB: not implemented"
	// The IsOnCurve function in Go's elliptic curve package
	// doesn't consider the point-at-infinity to be "on the curve"
	return false
}

// Try to generate a point on this curve from a chosen x-coordinate,
// with a random sign.
func (P *curvePoint) genPoint(x *big.Int, rand cipher.Stream) bool {
	_ = "STUB: not implemented"
	// Compute the corresponding Y coordinate, if any
	return false
}

// Pick a random sign for the y coordinate

// Check that it's a valid point

// Doesn't yield a valid point!

func (P *curvePoint) EmbedLen() int {
	_ = "STUB: not implemented"
	// Reserve at least 8 most-significant bits for randomness,
	// and the least-significant 8 bits for embedded data length.
	// (Hopefully it's unlikely we'll need >=2048-bit curves soon.)
	return 0
}

func (P *curvePoint) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Embed picks a curve point containing a variable amount of embedded data.
// Remaining bits comprising the point are chosen randomly.
func (P *curvePoint) Embed(data []byte, rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Encode length in low 8 bits
// Copy in data to embed

// Data extracts embedded data from a curve point
func (P *curvePoint) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// pad leading zero bytes if necessary

func (P *curvePoint) Add(A, B kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *
	//nolint:errcheck // Design pattern to emulate generics
	new(kyber.Point)
}

//nolint:errcheck // Design pattern to emulate generics

func (P *curvePoint) Sub(A, B kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *
	//nolint:errcheck // Design pattern to emulate generics
	new(kyber.Point)
}

//nolint:errcheck // Design pattern to emulate generics

//nolint:errcheck // Design pattern to emulate generics

func (P *curvePoint) Neg(A kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (P *curvePoint) Mul(s kyber.Scalar, B kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Point)
}

//nolint:errcheck // Design pattern to emulate generics

func (P *curvePoint) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

// uncompressed ANSI X9.62 representation

// MarshalBinary marshall this point into binary format according to
// SEC 1, Version 2.0, Section 2.3.3
func (P *curvePoint) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	// Note: explicit implementation since elliptic.Marshall is deprecated
	return nil, nil
}

// uncompressed point format

// UnmarshalBinary unmarshalls the given buffer into the receiver according
// to SEC 1, Version 2.0, Section 2.3.4
func (P *curvePoint) UnmarshalBinary(buf []byte) error {
	_ = "STUB: not implemented"
	// Note: explicit implementation since elliptic.Unmarshall is deprecated
	return nil
}

func (P *curvePoint) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (P *curvePoint) UnmarshalFrom(r io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// interface for curve-specifc mathematical functions
type curveOps interface {
	sqrt(y *big.Int) *big.Int
}

// Curve is an implementation of the kyber.Group interface
// for NIST elliptic curves, built on Go's native elliptic curve library.
type curve struct {
	elliptic.Curve
	curveOps
	p *elliptic.CurveParams
}

// Return the number of bytes in the encoding of a Scalar for this curve.
func (c *curve) ScalarLen() int { _ = "STUB: not implemented"; return 0 }

// Create a Scalar associated with this curve. The scalars created by
// this package implement kyber.Scalar's SetBytes method, interpreting
// the bytes as a big-endian integer, so as to be compatible with the
// Go standard library's big.Int type.
func (c *curve) Scalar() kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }

// Number of bytes required to store one coordinate on this curve
func (c *curve) coordLen() int { _ = "STUB: not implemented"; return 0 }

// Return the number of bytes in the encoding of a Point for this curve.
// Currently uses uncompressed ANSI X9.62 format with both X and Y coordinates;
// this could change.
func (c *curve) PointLen() int { _ = "STUB: not implemented"; return 0 }

// ANSI X9.62: 1 header byte plus 2 coords

// Create a Point associated with this curve.
func (c *curve) Point() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *curvePoint) Set(A kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (P *curvePoint) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

// Return the order of this curve: the prime N in the curve parameters.
func (c *curve) Order() *big.Int { _ = "STUB: not implemented"; return nil }
