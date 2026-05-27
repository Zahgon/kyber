//go:build !constantTime

package p256

import (
	"crypto/cipher"
	"hash"
	"io"
	"reflect"

	"go.dedis.ch/kyber/v4"
)

// QrSuite is a quadratic residue suite
type QrSuite struct {
	ResidueGroup
}

// Hash returns the instance associated with the suite
func (s QrSuite) Hash() hash.Hash {
	_ = "STUB: not implemented"
	return *

	// XOF creates the XOF associated with the suite
	new(hash.Hash)
}

func (s QrSuite) XOF(key []byte) kyber.XOF {
	_ = "STUB: not implemented"
	return *

	// RandomStream returns a cipher.Stream that returns a key stream
	// from crypto/rand.
	new(kyber.XOF)
}

func (s QrSuite) RandomStream() cipher.Stream {
	_ = "STUB: not implemented"
	return *new(cipher.Stream)
}

func (s *QrSuite) Read(r io.Reader, objs ...any) error { _ = "STUB: not implemented"; return nil }

func (s *QrSuite) Write(w io.Writer, objs ...any) error { _ = "STUB: not implemented"; return nil }

// New implements the kyber.encoding interface
func (s *QrSuite) New(t reflect.Type) any { _ = "STUB: not implemented"; return *new(any) }

// NewBlakeSHA256QR512 returns a cipher suite based on package
// go.dedis.ch/kyber/v4/xof/blake2xb, SHA-256, and a residue group of
// quadratic residues modulo a 512-bit prime.
//
// This group size should be used only for testing and experimentation.
// 512-bit DSA-style groups are no longer considered secure.
func NewBlakeSHA256QR512() *QrSuite {
	_ = "STUB: not implemented"
	//nolint:lll
	return nil
}

//nolint:lll
