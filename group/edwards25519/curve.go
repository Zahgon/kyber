package edwards25519

import (
	"crypto/cipher"

	"go.dedis.ch/kyber/v4"
)

// Curve represents the Ed25519 group.
// There are no parameters and no initialization is required
// because it supports only this one specific curve.
type Curve struct {
}

// Return the name of the curve, "Ed25519".
func (c *Curve) String() string {
	_ = "STUB: not implemented"

	// ScalarLen returns 32, the size in bytes of an encoded Scalar
	// for the Ed25519 curve.
	return ""
}

func (c *Curve) ScalarLen() int {
	_ = "STUB: not implemented"

	// Scalar creates a new Scalar for the prime-order subgroup of the Ed25519 curve.
	// The scalars in this package implement kyber.Scalar's SetBytes
	// method, interpreting the bytes as a little-endian integer, in order to remain
	// compatible with other Ed25519 implementations, and with the standard implementation
	// of the EdDSA signature.
	return 0
}

func (c *Curve) Scalar() kyber.Scalar {
	_ = "STUB: not implemented"

	// PointLen returns 32, the size in bytes of an encoded Point on the Ed25519 curve.
	return *new(kyber.Scalar)
}

func (c *Curve) PointLen() int {
	_ = "STUB: not implemented"

	// Point creates a new Point on the Ed25519 curve.
	return 0
}

func (c *Curve) Point() kyber.Point {
	_ = "STUB: not implemented"
	return *

	// NewKeyAndSeedWithInput returns a formatted Ed25519 key (avoid subgroup attack by
	// requiring it to be a multiple of 8). It also returns the input and the digest used
	// to generate the key.
	new(kyber.Point)
}

func (c *Curve) NewKeyAndSeedWithInput(buffer []byte) (kyber.Scalar, []byte, []byte) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), nil, nil
}

//nolint:errcheck // Design pattern to emulate generics

// NewKeyAndSeed returns a formatted Ed25519 key (avoid subgroup attack by requiring
// it to be a multiple of 8). It also returns the seed and the input used to generate
// the key.
func (c *Curve) NewKeyAndSeed(stream cipher.Stream) (kyber.Scalar, []byte, []byte) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), nil, nil
}

// NewKey returns a formatted Ed25519 key (avoiding subgroup attack by requiring
// it to be a multiple of 8). NewKey implements the kyber/util/key.Generator interface.
func (c *Curve) NewKey(stream cipher.Stream) kyber.Scalar {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar)
}
