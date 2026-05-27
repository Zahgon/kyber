package vss

import (
	"crypto/cipher"
	"hash"

	"go.dedis.ch/kyber/v4"
)

// dhExchange computes the shared key from a private key and a public key
func dhExchange(suite Suite, ownPrivate kyber.Scalar, remotePublic kyber.Point) kyber.Point {
	_ = "STUB: not implemented"
	return *new(kyber.Point)
}

var sharedKeyLength = 32

// newAEAD returns the AEAD cipher to be use to encrypt a share
func newAEAD(fn func() hash.Hash, preSharedKey kyber.Point, context []byte) (cipher.AEAD, error) {
	_ = "STUB: not implemented"
	return *new(cipher.AEAD), nil
}

// keySize is arbitrary, make it long enough to seed the XOF
const keySize = 128

// context returns the context slice to be used when encrypting a share
func context(suite Suite, dealer kyber.Point, verifiers []kyber.Point) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
