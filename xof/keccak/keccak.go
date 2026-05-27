// Package keccak provides an implementation of kyber.XOF based on the
// Shake256 hash.
package keccak

import (
	"go.dedis.ch/kyber/v4"
	"golang.org/x/crypto/sha3"
)

type xof struct {
	sh   sha3.ShakeHash
	seed []byte
	// key is here to not make excess garbage during repeated calls
	// to XORKeyStream.
	key []byte
}

// New creates a new XOF using the Shake256 hash.
func New(seed []byte) kyber.XOF { _ = "STUB: not implemented"; return *new(kyber.XOF) }

func (x *xof) Clone() kyber.XOF { _ = "STUB: not implemented"; return *new(kyber.XOF) }

func (x *xof) Reseed() { _ = "STUB: not implemented"; return }

func (x *xof) Reset() { _ = "STUB: not implemented"; return }

func (x *xof) Read(dst []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (x *xof) Write(src []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (x *xof) XORKeyStream(dst, src []byte) { _ = "STUB: not implemented"; return }
