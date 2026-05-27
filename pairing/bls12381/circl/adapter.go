package circl

import (
	"go.dedis.ch/kyber/v4"
)

// SuiteBLS12381 is an adapter that implements the suites.Suite interface so that
// bls12381 can be used as a common suite to generate key pairs for instance but
// still preserves the properties of the pairing (e.g. the Pair function).
//
// It's important to note that the Point function will generate a point
// compatible with public keys only (group G2) where the signature must be
// used as a point from the group G1.
type SuiteBLS12381 struct {
	Suite
}

var _ kyber.Group = (*SuiteBLS12381)(nil)

// NewSuiteBLS12381 makes a new BN256 suite
func NewSuiteBLS12381() *SuiteBLS12381 { _ = "STUB: not implemented"; return nil }

// Point generates a point from the G2 group that can only be used
// for public keys
func (s *SuiteBLS12381) Point() kyber.Point {
	_ = "STUB: not implemented"
	return *

	// PointLen returns the length of a G2 point
	new(kyber.Point)
}

func (s *SuiteBLS12381) PointLen() int { _ = "STUB: not implemented"; return 0 }

// Scalar generates a scalar
func (s *SuiteBLS12381) Scalar() kyber.Scalar {
	_ = "STUB: not implemented"
	return *

	// ScalarLen returns the length of a scalar
	new(kyber.Scalar)
}

func (s *SuiteBLS12381) ScalarLen() int { _ = "STUB: not implemented"; return 0 }

// String returns the name of the suite
func (s *SuiteBLS12381) String() string { _ = "STUB: not implemented"; return "" }
