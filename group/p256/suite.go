//go:build !constantTime

package p256

import (
	"crypto/cipher"
	"hash"
	"io"
	"reflect"

	"go.dedis.ch/kyber/v4"
)

// Suite128 is the suite for P256 curve
type Suite128 struct {
	p256
}

// Hash returns the instance associated with the suite
func (s *Suite128) Hash() hash.Hash {
	_ = "STUB: not implemented"
	return *

	// XOF creates the XOF associated with the suite
	new(hash.Hash)
}

func (s *Suite128) XOF(key []byte) kyber.XOF {
	_ = "STUB: not implemented"
	return *

	// RandomStream returns a cipher.Stream that returns a key stream
	// from crypto/rand.
	new(kyber.XOF)
}

func (s *Suite128) RandomStream() cipher.Stream {
	_ = "STUB: not implemented"
	return *new(cipher.Stream)
}

func (s *Suite128) Read(r io.Reader, objs ...any) error { _ = "STUB: not implemented"; return nil }

func (s *Suite128) Write(w io.Writer, objs ...any) error { _ = "STUB: not implemented"; return nil }

// New implements the kyber.encoding interface
func (s *Suite128) New(t reflect.Type) any { _ = "STUB: not implemented"; return *new(any) }

// NewBlakeSHA256P256 returns a cipher suite based on package
// go.dedis.ch/kyber/v4/xof/blake2xb, SHA-256, and the NIST P-256
// elliptic curve. It returns random streams from Go's crypto/rand.
//
// The scalars created by this group implement kyber.Scalar's SetBytes
// method, interpreting the bytes as a big-endian integer, so as to be
// compatible with the Go standard library's big.Int type.
func NewBlakeSHA256P256() *Suite128 { _ = "STUB: not implemented"; return nil }
