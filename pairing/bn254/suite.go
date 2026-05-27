//go:build !constantTime

// package bn254 implements a particular bilinear group.
//
// Bilinear groups are the basis of many of the new cryptographic protocols that
// have been proposed over the past decade. They consist of a triplet of groups
// (G₁, G₂ and GT) such that there exists a function e(g₁ˣ,g₂ʸ)=gTˣʸ (where gₓ
// is a generator of the respective group). That function is called a pairing
// function.
//
// This package specifically implements the Optimal Ate pairing over a 256-bit
// Barreto-Naehrig curve as described in
// http://cryptojedi.org/papers/dclxvi-20100714.pdf. Its output is compatible
// with the implementation described in that paper.
//
// This package previously claimed to operate at a 128-bit security level.
// However, recent improvements in attacks mean that is no longer true. See
// https://moderncrypto.org/mail-archive/curves/2016/000740.html.
package bn254

import (
	"crypto/cipher"
	"hash"
	"io"
	"reflect"

	"go.dedis.ch/kyber/v4"
)

// Suite implements the pairing.Suite interface for the BN254 bilinear pairing.
type Suite struct {
	*commonSuite
	g1 *groupG1
	g2 *groupG2
	gt *groupGT
}

func newDefaultDomainG1() []byte { _ = "STUB: not implemented"; return nil }

func newDefaultDomainG2() []byte { _ = "STUB: not implemented"; return nil }

// NewSuite generates and returns a new BN254 pairing suite.
func NewSuite() *Suite { _ = "STUB: not implemented"; return nil }

// NewSuiteG1 returns a G1 suite.
func NewSuiteG1() *Suite { _ = "STUB: not implemented"; return nil }

// NewSuiteG2 returns a G2 suite.
func NewSuiteG2() *Suite { _ = "STUB: not implemented"; return nil }

// NewSuiteGT returns a GT suite.
func NewSuiteGT() *Suite { _ = "STUB: not implemented"; return nil }

// NewSuiteRand generates and returns a new BN254 suite seeded by the
// given cipher stream.
func NewSuiteRand(rand cipher.Stream) *Suite { _ = "STUB: not implemented"; return nil }

// Set G1 DST
func (s *Suite) SetDomainG1(dst []byte) { _ = "STUB: not implemented"; return }

// Set G2 DST
func (s *Suite) SetDomainG2(dst []byte) { _ = "STUB: not implemented"; return }

// G1 returns the group G1 of the BN254 pairing.
func (s *Suite) G1() kyber.Group {
	_ = "STUB: not implemented"

	// G2 returns the group G2 of the BN254 pairing.
	return *new(kyber.Group)
}

func (s *Suite) G2() kyber.Group {
	_ = "STUB: not implemented"

	// GT returns the group GT of the BN254 pairing.
	return *new(kyber.Group)
}

func (s *Suite) GT() kyber.Group {
	_ = "STUB: not implemented"

	// Pair takes the points p1 and p2 in groups G1 and G2, respectively, as input
	// and computes their pairing in GT.
	return *new(kyber.Group)
}

func (s *Suite) Pair(p1 kyber.Point, p2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// NB: Not safe for concurrent calls
func (s *Suite) ValidatePairing(p1, p2, inv1, inv2 kyber.Point) bool {
	_ = "STUB: not implemented"
	return false
}

var tScalar = reflect.TypeFor[kyber.Scalar]()
var tPoint = reflect.TypeFor[kyber.Point]()
var tPointG1 = reflect.TypeFor[pointG1]()
var tPointG2 = reflect.TypeFor[pointG2]()
var tPointGT = reflect.TypeFor[pointGT]()

type commonSuite struct {
	s cipher.Stream
	// kyber.Group is only set if we have a combined Suite
	kyber.Group
}

// New implements the kyber.Encoding interface.
func (c *commonSuite) New(t reflect.Type) any { _ = "STUB: not implemented"; return *new(any) }

// Read is the default implementation of kyber.Encoding interface Read.
func (c *commonSuite) Read(r io.Reader, objs ...any) error { _ = "STUB: not implemented"; return nil }

// Write is the default implementation of kyber.Encoding interface Write.
func (c *commonSuite) Write(w io.Writer, objs ...any) error { _ = "STUB: not implemented"; return nil }

// Hash returns a newly instantiated keccak256 hash function.
func (c *commonSuite) Hash() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

// XOF returns a newlly instantiated blake2xb XOF function.
func (c *commonSuite) XOF(seed []byte) kyber.XOF {
	_ = "STUB: not implemented"
	return *

	// RandomStream returns a cipher.Stream which corresponds to a key stream from
	// crypto/rand.
	new(kyber.XOF)
}

func (c *commonSuite) RandomStream() cipher.Stream {
	_ = "STUB: not implemented"
	return *new(cipher.Stream)
}

// String returns a recognizable string that this is a combined suite.
func (c commonSuite) String() string { _ = "STUB: not implemented"; return "" }
