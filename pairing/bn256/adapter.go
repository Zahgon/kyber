//go:build !constantTime

package bn256

import (
	"go.dedis.ch/kyber/v4"
)

// SuiteBn256 is an adapter that implements the suites.Suite interface so that
// bn256 can be used as a common suite to generate key pairs for instance but
// still preserves the properties of the pairing (e.g. the Pair function).
//
// It's important to note that the Point function will generate a point
// compatible with public keys only (group G2) where the signature must be
// used as a point from the group G1.
type SuiteBn256 struct {
	*Suite
}

var _ kyber.Group = (*SuiteBn256)(nil)

// NewSuiteBn256 makes a new BN256 suite
func NewSuiteBn256() *SuiteBn256 { _ = "STUB: not implemented"; return nil }

// Point generates a point from the G2 group that can only be used
// for public keys
func (s *SuiteBn256) Point() kyber.Point {
	_ = "STUB: not implemented"
	return *

	// PointLen returns the length of a G2 point
	new(kyber.Point)
}

func (s *SuiteBn256) PointLen() int { _ = "STUB: not implemented"; return 0 }

// Scalar generates a scalar
func (s *SuiteBn256) Scalar() kyber.Scalar {
	_ = "STUB: not implemented"
	return *

	// ScalarLen returns the lentgh of a scalar
	new(kyber.Scalar)
}

func (s *SuiteBn256) ScalarLen() int { _ = "STUB: not implemented"; return 0 }

// String returns the name of the suite
func (s *SuiteBn256) String() string { _ = "STUB: not implemented"; return "" }
