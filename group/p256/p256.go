//go:build !constantTime

package p256

import (
	"math/big"
)

// P256 implements the kyber.Group interface
// for the NIST P-256 elliptic curve,
// based on Go's native elliptic curve library.
type p256 struct {
	curve
}

func (curve *p256) String() string {
	_ = "STUB: not implemented"

	// Optimized modular square root for P-256 curve, from
	// "Mathematical routines for the NIST prime elliptic curves" (April 2010)
	//
	//nolint:gocritic // not actually comments, help understand the code
	return ""
}

func (curve *p256) sqrt(c *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }

// t1 = c^(2^2-1)

// t2 = c^(2^4-1)

// t3 = c^(2^8-1)

// t4 = c^(2^16-1)

// r = c^(2^32-1)

// r = c^(2^64-2^32+1)

// r = c^(2^160-2^128+2^96+1)

// r = c^(2^254-2^222+2^190+2^94) = sqrt(c) mod p256

// Init initializes standard Curve instances
func (curve *p256) Init() curve { _ = "STUB: not implemented"; return *new(curve) }
