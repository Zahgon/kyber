package edwards25519

import (
	"crypto/cipher"
	"hash"
	"io"
	"reflect"

	"go.dedis.ch/kyber/v4"
)

// SuiteEd25519 implements some basic functionalities such as Group, HashFactory,
// and XOFFactory.
type SuiteEd25519 struct {
	Curve

	r cipher.Stream
}

// Hash returns a newly instanciated sha256 hash function.
func (s *SuiteEd25519) Hash() hash.Hash {
	_ = "STUB: not implemented"
	return *

	// XOF returns an XOF which is implemented via the Blake2b hash.
	new(hash.Hash)
}

func (s *SuiteEd25519) XOF(key []byte) kyber.XOF { _ = "STUB: not implemented"; return *new(kyber.XOF) }

func (s *SuiteEd25519) Read(r io.Reader, objs ...any) error { _ = "STUB: not implemented"; return nil }

func (s *SuiteEd25519) Write(w io.Writer, objs ...any) error { _ = "STUB: not implemented"; return nil }

// New implements the kyber.Encoding interface
func (s *SuiteEd25519) New(t reflect.Type) any { _ = "STUB: not implemented"; return *new(any) }

// RandomStream returns a cipher.Stream that returns a key stream
// from crypto/rand.
func (s *SuiteEd25519) RandomStream() cipher.Stream {
	_ = "STUB: not implemented"
	return *new(cipher.Stream)
}

// NewBlakeSHA256Ed25519 returns a cipher suite based on package
// go.dedis.ch/kyber/v4/xof/blake2xb, SHA-256, and the Ed25519 curve.
// It produces cryptographically random numbers via package crypto/rand.
func NewBlakeSHA256Ed25519() *SuiteEd25519 { _ = "STUB: not implemented"; return nil }

// NewBlakeSHA256Ed25519WithRand returns a cipher suite based on package
// go.dedis.ch/kyber/v4/xof/blake2xb, SHA-256, and the Ed25519 curve.
// It produces cryptographically random numbers via the provided stream r.
func NewBlakeSHA256Ed25519WithRand(r cipher.Stream) *SuiteEd25519 {
	_ = "STUB: not implemented"
	return nil
}
