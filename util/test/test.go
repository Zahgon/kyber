package test

import (
	"crypto/cipher"
	"testing"

	"go.dedis.ch/kyber/v4"
)

// Suite represents the functionalities that this package can test
type suite interface {
	kyber.Group
	kyber.HashFactory
	kyber.XOFFactory
	kyber.Random
}

type suiteStable struct {
	suite
	xof kyber.XOF
}

func newSuiteStable(s suite) *suiteStable { _ = "STUB: not implemented"; return nil }

func (ss *suiteStable) RandomStream() cipher.Stream {
	_ = "STUB: not implemented"
	return *new(cipher.Stream)
}

func testEmbed(t *testing.T, g kyber.Group, rand cipher.Stream, points *[]kyber.Point,
	s string) {
	_ = "STUB: not implemented"
	return
}

func testPointSet(t *testing.T, g kyber.Group, rand cipher.Stream) {
	_ = "STUB: not implemented"
	return
}

func testPointClone(t *testing.T, g kyber.Group, rand cipher.Stream) {
	_ = "STUB: not implemented"
	return
}

func testScalarSet(t *testing.T, g kyber.Group, rand cipher.Stream) {
	_ = "STUB: not implemented"
	return
}

func testScalarClone(t *testing.T, g kyber.Group, rand cipher.Stream) {
	_ = "STUB: not implemented"
	return
}

func testSanityCheck(
	t *testing.T,
	points []kyber.Point,
	g kyber.Group,
	stmp, s1, s2 kyber.Scalar,
	gen, ptmp kyber.Point,
) ([]kyber.Point, kyber.Point, kyber.Point, kyber.Point, bool) {
	_ = "STUB: not implemented"
	// Sanity-check relationship between addition and multiplication
	return nil, *new(kyber.Point), *new(kyber.Point), *new(kyber.Point), false
}

// Find out if this curve has a prime order:
// if the curve does not offer a method IsPrimeOrder,
// then assume that it is.

// Verify additive and multiplicative identities of the generator.

// secret.Inv works only in prime-order groups

func testHomomorphicIdentities(
	t *testing.T,
	primeOrder bool,
	g kyber.Group,
	gen, ptmp, p1, p2, dh1 kyber.Point,
	stmp, s1, s2 kyber.Scalar) {
	_ = "STUB: not implemented"
	// Additive homomorphic identities
	return
}

// Multiplicative homomorphic identities

func testRandomlyPickedPoint(
	t *testing.T,
	primeOrder bool,
	points []kyber.Point,
	g kyber.Group,
	gen, ptmp kyber.Point,
	stmp kyber.Scalar,
	rand cipher.Stream,
) []kyber.Point {
	_ = "STUB: not implemented"
	return nil
}

func testEncodingDecoding(t *testing.T, g kyber.Group, ptmp kyber.Point, stmp kyber.Scalar, rand cipher.Stream) {
	_ = "STUB: not implemented"
	return
}

// Apply a generic set of validation tests to a cryptographic Group,
// using a given source of [pseudo-]randomness.
//
// Returns a log of the pseudorandom Points produced in the test,
// for comparison across alternative implementations
// that are supposed to be equivalent.
func testGroup(t *testing.T, g kyber.Group, rand cipher.Stream) []kyber.Point {
	_ = "STUB: not implemented"
	return nil
}

// Do a simple Diffie-Hellman test

// Sanity-check relationship between addition and multiplication

// Test secret inverse to get from dh1 back to p1

// Zero and One identity secrets

// homomorphic identities

// Test randomly picked points

// Test embedding data

// Test verifiable secret sharing
// Test encoding and decoding

// Test that we can marshal/ unmarshal null point

// GroupTest applies a generic set of validation tests to a cryptographic Group.
func GroupTest(t *testing.T, g kyber.Group) { _ = "STUB: not implemented"; return }

// CompareGroups tests two group implementations that are supposed to be equivalent,
// and compare their results.
func CompareGroups(t *testing.T, fn func(key []byte) kyber.XOF, g1, g2 kyber.Group) {
	_ = "STUB: not implemented"

	// Produce test results from the same pseudorandom seed
	return
}

// Compare resulting Points

// SuiteTest tests a standard set of validation tests to a ciphersuite.
func SuiteTest(t *testing.T, suite suite) {
	_ = "STUB: not implemented"

	// Try hashing something
	return
}

// Generate some pseudorandom bits

// Test if it generates two fresh keys

// Test if it creates the same key with the same seed

// Test the public-key group arithmetic
