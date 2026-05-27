//go:build !constantTime

package bn256

import "go.dedis.ch/kyber/v4"

// HashG1 implements a hashing function into the G1 group.
//
// dst represents domain separation tag, similar to salt, for the hash.
func HashG1(msg, dst []byte) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

//nolint:funlen
func mapToCurve(t *gfP) kyber.Point {
	_ = "STUB: not implemented"

	// calculate w = (s * t)/(1 + B + t^2)
	// we calculate w0 = s * t * (1 + B + t^2) and inverse of it, so that w = (st)^2/w0
	// and then later x3 = 1 + (1 + B + t^2)^4/w0^2
	return *new(kyber.Point)
}

// Computing a = (1 + B + t^2)

// calculate x1 = ((-1 + s) / 2) - t * w

// check if y=x1^3+3 is a square

// calculate x2 = -1 - x1

// check if y=x2^3+3 is a square

// calculate x3 = 1 + (1/ww) = 1 + a^4 * w0^2
