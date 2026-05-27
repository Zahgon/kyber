//go:build !constantTime

package kilic

import (
	"crypto/cipher"
	"hash"
	"io"
	"reflect"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
)

type Suite struct {
	domainG1 []byte
	domainG2 []byte
}

// NewBLS12381Suite is the same as calling NewBLS12381SuiteWithDST(nil, nil): it uses the default domain separation
// tags for its Hash To Curve functions.
func NewBLS12381Suite() pairing.Suite {
	_ = "STUB: not implemented"

	// NewBLS12381SuiteWithDST allows you to set your own domain separation tags to be used by the Hash To Curve functions.
	// Since the DST shouldn't be 0 len, if you provide nil or a 0 len byte array, it will use the RFC default values.
	return *new(pairing.Suite)
}

func NewBLS12381SuiteWithDST(DomainG1, DomainG2 []byte) pairing.Suite {
	_ = "STUB: not implemented"
	return *new(pairing.Suite)
}

func (s *Suite) SetDomainG1(dst []byte) { _ = "STUB: not implemented"; return }

func (s *Suite) G1() kyber.Group { _ = "STUB: not implemented"; return *new(kyber.Group) }

func (s *Suite) SetDomainG2(dst []byte) { _ = "STUB: not implemented"; return }

func (s *Suite) G2() kyber.Group { _ = "STUB: not implemented"; return *new(kyber.Group) }

func (s *Suite) GT() kyber.Group {
	_ = "STUB: not implemented"
	return *

	// ValidatePairing implements the `pairing.Suite` interface
	new(kyber.Group)
}

func (s *Suite) ValidatePairing(p1, p2, p3, p4 kyber.Point) bool {
	_ = "STUB: not implemented"
	return false

	// we need to clone the point because of https://github.com/kilic/bls12-381/issues/37
	// in order to avoid risks of race conditions.
}

func (s *Suite) Pair(p1, p2 kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

// New implements the kyber.Encoding interface.
func (s *Suite) New(_ reflect.Type) any { _ = "STUB: not implemented"; return *new(any) }

// Read is the default implementation of kyber.Encoding interface Read.
func (s *Suite) Read(_ io.Reader, _ ...any) error { _ = "STUB: not implemented"; return nil }

// Write is the default implementation of kyber.Encoding interface Write.
func (s *Suite) Write(_ io.Writer, _ ...any) error { _ = "STUB: not implemented"; return nil }

// Hash returns a newly instantiated sha256 hash function.
func (s *Suite) Hash() hash.Hash {
	_ = "STUB: not implemented"
	return *

	// XOF returns a newly instantiated blake2xb XOF function.
	new(hash.Hash)
}

func (s *Suite) XOF(seed []byte) kyber.XOF {
	_ = "STUB: not implemented"
	return *

	// RandomStream returns a cipher.Stream which corresponds to a key stream from
	// crypto/rand.
	new(kyber.XOF)
}

func (s *Suite) RandomStream() cipher.Stream { _ = "STUB: not implemented"; return *new(cipher.Stream) }
