//go:build !constantTime

package p256

import (
	"crypto/cipher"
	"errors"
	"io"
	"math/big"

	"go.dedis.ch/kyber/v4"
)

var ErrTypeCast = errors.New("invalid type cast")

var one = big.NewInt(1)
var two = big.NewInt(2)

type residuePoint struct {
	big.Int
	g *ResidueGroup
}

// Steal value from DSA, which uses recommendation from FIPS 186-3
const numMRTests = 64

// Probabilistically test whether a big integer is prime.
func isPrime(i *big.Int) bool { _ = "STUB: not implemented"; return false }

func (P *residuePoint) String() string { _ = "STUB: not implemented"; return "" }

func (P *residuePoint) Equal(p2 kyber.Point) bool { _ = "STUB: not implemented"; return false }

func (P *residuePoint) Null() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *residuePoint) Base() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *residuePoint) Set(P2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (P *residuePoint) Clone() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

func (P *residuePoint) Valid() bool { _ = "STUB: not implemented"; return false }

func (P *residuePoint) EmbedLen() int {
	_ = "STUB: not implemented"
	// Reserve at least 8 most-significant bits for randomness,
	// and the least-significant 16 bits for embedded data length.
	return 0
}

func (P *residuePoint) Pick(rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Embed the given data with some pseudo-random bits.
// This will only work efficiently for quadratic residue groups!
func (P *residuePoint) Embed(data []byte, rand cipher.Stream) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// Encode length in low 16 bits

// Copy in embedded data

// Extract embedded data from a Residue group element
func (P *residuePoint) Data() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// pad leading zero bytes if necessary

func (P *residuePoint) Add(A, B kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (P *residuePoint) Sub(A, B kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (P *residuePoint) Neg(A kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

func (P *residuePoint) Mul(s kyber.Scalar, B kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// to protect against golang/go#22830

func (P *residuePoint) MarshalSize() int { _ = "STUB: not implemented"; return 0 }

func (P *residuePoint) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	// may be shorter than len(buf)
	return nil, nil
}

func (P *residuePoint) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (P *residuePoint) MarshalTo(w io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (P *residuePoint) UnmarshalFrom(r io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Parameters represents the domain parameters for a key. These parameters can
// be shared across many keys. The bit length of Q must be a multiple of 8.
// This struct is copy-pasted directly from crypto/dsa since it is deprecated,
// and we want to avoid dependencies on it.
type Parameters struct {
	P, Q, G *big.Int
}

/*
A ResidueGroup represents a DSA-style modular integer arithmetic group,
defined by two primes P and Q and an integer R, such that P = Q*R+1.
Points in a ResidueGroup are R-residues modulo P,
and Scalars are integer exponents modulo the group order Q.

In traditional DSA groups P is typically much larger than Q,
and hence use a large multiple R.
This is done to minimize the computational cost of modular exponentiation
while maximizing security against known classes of attacks:
P must be on the order of thousands of bits long
while for security Q is believed to require only hundreds of bits.
Such computation-optimized groups are suitable
for Diffie-Hellman agreement, DSA or ElGamal signatures, etc.,
which depend on Point.Mul() and homomorphic properties.

However, residue groups with large R are less suitable for
public-key cryptographic techniques that require choosing Points
pseudo-randomly or to contain embedded data,
as required by ElGamal encryption for example.
For such purposes quadratic residue groups are more suitable -
representing the special case where R=2 and hence P=2Q+1.
As a result, the Point.Pick() method should be expected to work efficiently
ONLY on quadratic residue groups in which R=2.
*/
type ResidueGroup struct {
	Parameters
	R *big.Int
}

func (g *ResidueGroup) String() string { _ = "STUB: not implemented"; return "" }

// ScalarLen returns the number of bytes in the encoding of a Scalar
// for this Residue group.
func (g *ResidueGroup) ScalarLen() int { _ = "STUB: not implemented"; return 0 }

// Scalar creates a Scalar associated with this Residue group,
// with an initial value of nil.
func (g *ResidueGroup) Scalar() kyber.Scalar { _ = "STUB: not implemented"; return *new(kyber.Scalar) }

// PointLen returns the number of bytes in the encoding of a Point
// for this Residue group.
func (g *ResidueGroup) PointLen() int { _ = "STUB: not implemented"; return 0 }

// Point creates a Point associated with this Residue group,
// with an initial value of nil.
func (g *ResidueGroup) Point() kyber.Point { _ = "STUB: not implemented"; return *new(kyber.Point) }

// Order returns the order of this Residue group, namely the prime Q.
func (g *ResidueGroup) Order() *big.Int {
	_ = "STUB: not implemented"

	// Valid validates the parameters for a Residue group,
	// checking that P and Q are prime, P=Q*R+1,
	// and that G is a valid generator for this group.
	return nil
}

func (g *ResidueGroup) Valid() bool {
	_ = "STUB: not implemented"

	// Make sure both P and Q are prime
	return false
}

// Validate the equation P = QR+1

// Validate the generator G

// SetParams explicitly initializes a ResidueGroup with given parameters.
func (g *ResidueGroup) SetParams(p, q, r, g1 *big.Int) { _ = "STUB: not implemented"; return }

// QuadraticResidueGroup initializes Residue group parameters for a quadratic residue group,
// by picking primes P and Q such that P=2Q+1
// and the smallest valid generator G for this group.
func (g *ResidueGroup) QuadraticResidueGroup(bitlen uint, rand cipher.Stream) {
	_ = "STUB: not implemented"

	// pick primes p,q such that p = 2q+1
	return
}

// First pick a prime Q

// must be odd

// TODO:Does the corresponding P come out prime too?

// pick standard generator G
