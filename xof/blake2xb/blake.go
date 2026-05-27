// Package blake2xb provides an implementation of kyber.XOF based on the
// Blake2xb construction.
package blake2xb

import (
	"go.dedis.ch/kyber/v4"
	"golang.org/x/crypto/blake2b"
)

type xof struct {
	impl blake2b.XOF
	seed []byte
	// key is here to not make excess garbage during repeated calls
	// to XORKeyStream.
	key []byte
}

// New creates a new XOF using the Blake2b hash.
func New(seed []byte) kyber.XOF { _ = "STUB: not implemented"; return *new(kyber.XOF) }

func (x *xof) Clone() kyber.XOF { _ = "STUB: not implemented"; return *new(kyber.XOF) }

func (x *xof) Read(dst []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (x *xof) Write(src []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (x *xof) Reseed() {
	_ = "STUB: not implemented"
	// Use New to create a new one seeded with output from the old one.
	return
}

// Steal the XOF implementation, and put it inside of x.

func (x *xof) Reset() { _ = "STUB: not implemented"; return }

func (x *xof) XORKeyStream(dst, src []byte) { _ = "STUB: not implemented"; return }
