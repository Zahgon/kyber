//go:build !constantTime

package edwards25519vartime

import (
	"crypto/cipher"
	"hash"
	"io"
	"reflect"

	"go.dedis.ch/kyber/v4"
)

// SuiteEd25519 is the suite for the Ed25519 curve
type SuiteEd25519 struct {
	ProjectiveCurve
}

// Hash returns the instance associated with the suite
func (s *SuiteEd25519) Hash() hash.Hash {
	_ = "STUB: not implemented"
	return *

	// XOF creates the XOF associated with the suite
	new(hash.Hash)
}

func (s *SuiteEd25519) XOF(seed []byte) kyber.XOF {
	_ = "STUB: not implemented"
	return *new(kyber.XOF)
}

func (s *SuiteEd25519) Read(r io.Reader, objs ...any) error { _ = "STUB: not implemented"; return nil }

func (s *SuiteEd25519) Write(w io.Writer, objs ...any) error { _ = "STUB: not implemented"; return nil }

// New implements the kyber.encoding interface
func (s *SuiteEd25519) New(t reflect.Type) any { _ = "STUB: not implemented"; return *new(any) }

// RandomStream returns a cipher.Stream that returns a key stream
// from crypto/rand.
func (s *SuiteEd25519) RandomStream() cipher.Stream {
	_ = "STUB: not implemented"
	return *

	// NewBlakeSHA256Ed25519 returns a cipher suite based on package
	// go.dedis.ch/kyber/v4/xof/blake2xb, SHA-256, and Ed25519.
	//
	// If fullGroup is false, then the group is the prime-order subgroup.
	//
	// The scalars created by this group implement kyber.Scalar's SetBytes
	// method, interpreting the bytes as a big-endian integer, so as to be
	// compatible with the Go standard library's compatible.Int type.
	new(cipher.Stream)
}

func NewBlakeSHA256Ed25519(fullGroup bool) *SuiteEd25519 { _ = "STUB: not implemented"; return nil }
