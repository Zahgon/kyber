// Package edwards25519 provides an optimized Go implementation of a
// Twisted Edwards curve that is isomorphic to Curve25519. For details see:
// http://ed25519.cr.yp.to/.
//
// This code is based on Adam Langley's Go port of the public domain,
// "ref10" implementation of the ed25519 signing scheme in C from SUPERCOP.
// It was generalized and extended to support full kyber.Group arithmetic
// by the DEDIS lab at Yale and EPFL.
//
// Due to the field element and group arithmetic optimizations
// described in the Ed25519 paper, this implementation generally
// performs extremely well, typically comparable to native C
// implementations.  The tradeoff is that this code is completely
// specialized to a single curve.
package edwards25519

import (
	"crypto/cipher"
	"hash"
	"io"

	"go.dedis.ch/kyber/v4"
	"golang.org/x/crypto/sha3"
)

var marshalPointID = [8]byte{'e', 'd', '.', 'p', 'o', 'i', 'n', 't'}
var longDomainSeparator = "H2C-OVERSIZE-DST-"

type point struct {
	ge      extendedGroupElement
	varTime bool
}

func (P *point) String() string { _ = "STUB: not implemented"; return "" }

func (P *point) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (P *point) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalID returns the type tag used in encoding/decoding
func (P *point) MarshalID() [8]byte { _ = "STUB: not implemented"; return nil }

func (P *point) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func (P *point) MarshalTo(w io.Writer) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (P *point) UnmarshalFrom(r io.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Equality test for two Points on the same curve
func (P *point) Equal(P2 kyber.Point) bool { _ = "STUB: not implemented"; return false }

// Set point to be equal to P2.
func (P *point) Set(P2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Set point to be equal to P2.
func (P *point) Clone() kyber.Point {
	_ = "STUB: not implemented"
	return *

	// Set to the neutral element, which is (0,1) for twisted Edwards curves.
	new(kyber.Point)
}

func (P *point) Null() kyber.Point {
	_ = "STUB: not implemented"
	return *

	// Set to the standard base point for this curve
	new(kyber.Point)
}

func (P *point) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *point) EmbedLen() int {
	_ = "STUB: not implemented"
	// Reserve the most-significant 8 bits for pseudo-randomness.
	// Reserve the least-significant 8 bits for embedded data length.
	// (Hopefully it's unlikely we'll need >=2048-bit curves soon.)
	return 0
}

func (P *point) Embed(data []byte, rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"

	// How many bytes to embed?
	return *new(kyber.Point)
}

// Pick a random point, with optional embedded data

// Encode length in low 8 bits
// Copy in data to embed

// Try to decode
// invalid point, retry

// If we're using the full group,
// we just need any point on the curve, so we're done.
//		if c.full {
//			return P,data[dl:]
//		}

// We're using the prime-order subgroup,
// so we need to make sure the point is in that subencoding.
// If we're not trying to embed data,
// we can convert our point into one in the subgroup
// simply by multiplying it by the cofactor.

// multiply by cofactor

// unlucky; try again

// success

// Since we need the point's y-coordinate to hold our data,
// we must simply check if the point is in the subgroup
// and retry point generation until it is.

// success

// Keep trying...

func (P *point) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Extract embedded data from a point group element
func (P *point) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// extract length byte

func (P *point) Add(P1, P2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Point)
}

//nolint:errcheck // Design pattern to emulate generics

func (P *point) Sub(P1, P2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	//nolint:errcheck // Design pattern to emulate generics
	return *new(kyber.Point)
}

//nolint:errcheck // Design pattern to emulate generics

// Neg finds the negative of point A.
// For Edwards curves, the negative of (x,y) is (-x,y).
func (P *point) Neg(A kyber.Point) kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

// Mul multiplies point p by scalar s using the repeated doubling method.
func (P *point) Mul(s kyber.Scalar, A kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// HasSmallOrder determines whether the group element has small order
//
// Provides resilience against malicious key substitution attacks (M-S-UEO)
// and message bound security (MSB) even for malicious keys
// See paper https://eprint.iacr.org/2020/823.pdf for definitions and theorems
//
// This is the same code as in
// https://github.com/jedisct1/libsodium/blob/4744636721d2e420f8bbe2d563f31b1f5e682229/src/libsodium/crypto_core/ed25519/ref10/ed25519_ref10.c#L1170
//
//nolint:lll // Url above
func (P *point) HasSmallOrder() bool { _ = "STUB: not implemented"; return false }

// Constant time verification if one or more of the c's are zero

// IsCanonical determines whether the group element is canonical
//
// Checks whether group element s is less than p, according to RFC8032§5.1.3.1
// https://tools.ietf.org/html/rfc8032#section-5.1.3
//
// Taken from
// https://github.com/jedisct1/libsodium/blob/4744636721d2e420f8bbe2d563f31b1f5e682229/src/libsodium/crypto_core/ed25519/ref10/ed25519_ref10.c#L1113
//
// The method accepts a buffer instead of calling `MarshalBinary` on the receiver
// because that always returns a value modulo `prime`.
//
//nolint:lll // Url above
func (P *point) IsCanonical(s []byte) bool { _ = "STUB: not implemented"; return false }

// subtraction might underflow

func (P *point) Hash(m []byte, dst string) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Clear cofactor

func hashToField(m []byte, dst string, count uint64) []fieldElement {
	_ = "STUB: not implemented"
	// L param in RFC9380 section 5
	// https://datatracker.ietf.org/doc/html/rfc9380#name-hashing-to-a-finite-field
	return nil
}

// todo, what's this?
// 	tv := compatible.NewInt(0).SetBytes(uniformBytes[elmOffset:elmOffset+l], prime)
// says that prime has a smaller size than l
// should we fix the modulus to be  1 << l ? Is it fine to pass through big.Int?

func expandMessageXMD(h hash.Hash, m []byte, domainSeparator string, byteLen uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Compute mPrime = Z_pad || msg || l_i_b_str || I2OSP(0, 1) || DST_prim

// Compute b0 = H(msg_prime)

// Compute b_1 = H(b_0 || I2OSP(1, 1) || DST_prime)

func expandMessageXOF(h sha3.ShakeHash, m []byte, domainSeparator string, byteLen uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// i2OSP converts a nonnegative integer to a byte array of a
// specified length. Implementation from [RFC8017]
func i2OSP(x uint64, xLen int32) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// create modulus int as the biggest value representable on xLen bytes

// Use the modulus to get the bytes of x

func byteXor(dst, b1, b2 []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// curve25519Elligator2 implements a map from fieldElement to a point on Curve25519
// as defined in section G.2.1. of [RFC9380]
// [RFC9380]: https://datatracker.ietf.org/doc/html/rfc9380#ell2-opt
//
//nolint:funlen
func curve25519Elligator2(u fieldElement) (xn, xd, yn, yd fieldElement) {
	_ = "STUB: not implemented"
	// Some const needed
	return *new(fieldElement), *new(fieldElement), *new(fieldElement), *new(fieldElement)
}

// c1 = (q + 3) / 8
// c2 = 2^c1
// c3 = sqrt(-1)
// c4 = (q - 5) / 8
// Computed with sagemath

// Temporary variables

// Compute tv1 = 2 * u^2
// Compute xd = 1 + tv1
// Compute x1n = -J
// Compute tv2 = xd^2
// Compute gxd = tv2 * xd
// Compute gx1 = J * tv1
// Compute gx1 = gx1 * x1n
// Compute gx1 = gx1 + tv2
// Compute gx1 = gx1 * x1n
// Compute tv3 = gxd^2
// Compute tv2 = tv3^2
// Compute tv3 = tv3 * gxd
// Compute tv3 = tv3 * gx1
// Compute tv2 = tv2 * tv3

// compute y11 = tv2 ^ c4

// Compute y11 = y11 * tv3
// Compute y12 = y11 * c3
// Compute tv2 = y11^2
// Compute tv2 = tv2 * gxd

// y1 = y11 if e1 == 1 else y12

// Compute x2n = x1n * tv1
// Compute y21 = y11 * u
// Compute y21 = y21 * c2
// Compute y22 = y21 * c3
// Compute gx2 = gx1 * tv1
// Compute tv2 = y21^2
// Compute tv2 = tv2 * gxd

// y2 = y21 if e == 1 else y22

// Compute tv2 = y1^2
// Compute tv2 = tv2 * gxd

// xn = x1n if e3 == 1 else x2n

// y = y1 if e4 == 1 else y2

// yNeg = -y
// y = yNeg if e3 XOR e4 == 1 else y

// mapToCurveElligator2Ed25519 implements a map from fieldElement to a point on ed25519
// as defined in section G.2.2. of [RFC9380]
// [RFC9380]: https://datatracker.ietf.org/doc/html/rfc9380#ell2-opt
func mapToCurveElligator2Ed25519(u fieldElement) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// c = sqrt(-486664)
// computed using sagemath

// Compute xn = xMn * yMd
// Compute xn = xn * c
// Compute xd = xMd * yMn
// Compute yn = xMn - xMd
// Compute yd = xMn + xMd
// Compute tv1 = xd * yd

// xn = 0 if e == 1 else xn
// xd = 1 if e == 1 else xd
// yn = 1 if e == 1 else yn
// yd = 1 if e == 1 else yd
