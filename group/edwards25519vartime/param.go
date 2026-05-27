//go:build !constantTime

// Package edwards25519vartime contains several implementations of Twisted Edwards Curves,
// from general and unoptimized to highly specialized and optimized.
//
// Twisted Edwards curves are elliptic curves satisfying the equation:
//
//	ax^2 + y^2 = c^2(1 + dx^2y^2)
//
// for some scalars c, d over some field K. We assume K is a (finite) prime field for a
// large prime p. We also assume c == 1 because all curves in the generalized form
// are isomorphic to curves having c == 1.
//
// For details see Bernstein et al, "Twisted Edwards Curves", http://eprint.iacr.org/2008/013.pdf
package edwards25519vartime

import (
	"go.dedis.ch/kyber/v4/compatible"
)

// Param defines a Twisted Edwards curve (TEC).
type Param struct {
	Name string // Name of curve

	P compatible.Int // Prime defining the underlying field
	Q compatible.Int // Order of the prime-order base point
	R int            // Cofactor: Q*R is the total size of the curve

	A, D compatible.Int // Edwards curve equation parameters

	FBX, FBY compatible.Int // Standard base point for full group
	PBX, PBY compatible.Int // Standard base point for prime-order subgroup

	Elligator1s compatible.Int // Optional s parameter for Elligator 1
	Elligator2u compatible.Int // Optional u parameter for Elligator 2
}

// Return the name of this curve.
func (p *Param) String() string {
	_ = "STUB: not implemented"

	// Param1174 defines Curve1174, as specified in:
	// Bernstein et al, "Elligator: Elliptic-curve points indistinguishable
	// from uniform random strings"
	// http://elligator.cr.yp.to/elligator-20130828.pdf
	return ""
}

func Param1174() *Param { _ = "STUB: not implemented"; return nil }

// todo what's the modulus here?

// Full-group generator is (4/V,3/5)

// Elligator1 parameter s for Curve1174 (Elligator paper section 4.1)

// ParamEd25519 defines the Edwards version of Curve25519, as specified in:
// Bernstein et al, "High-speed high-security signatures",
// http://ed25519.cr.yp.to/ed25519-20110926.pdf
func ParamEd25519() *Param { _ = "STUB: not implemented"; return nil }

// Non-square u for Elligator2

// ParamE382 defines the E-382 curve specified in:
// Aranha et al, "A note on high-security general-purpose elliptic curves",
// http://eprint.iacr.org/2013/647.pdf
//
// and more recently in:
//
// "Additional Elliptic Curves for IETF protocols"
// http://tools.ietf.org/html/draft-ladd-safecurves-02
// (this I-D is now expired)
func ParamE382() *Param { _ = "STUB: not implemented"; return nil }

// p = 2^382-105

//nolint:lll // Line not breakable

// Param41417 defines the Curve41417 curve, as specified in:
// Bernstein et al, "Curve41417: Karatsuba revisited",
// http://eprint.iacr.org/2014/526.pdf
func Param41417() *Param { _ = "STUB: not implemented"; return nil }

//nolint:lll // Line not breakable

// ParamE521 defines the E-521 curve specified in:
// Aranha et al, "A note on high-security general-purpose elliptic curves",
// http://eprint.iacr.org/2013/647.pdf
//
// and more recently included in:
// "Additional Elliptic Curves for IETF protocols"
// http://tools.ietf.org/html/draft-ladd-safecurves-02
func ParamE521() *Param { _ = "STUB: not implemented"; return nil }

//nolint:lll // Line not breakable
