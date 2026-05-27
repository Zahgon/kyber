// Package eddsa implements the EdDSA signature algorithm according to
// RFC8032.
package eddsa

import (
	"crypto/cipher"
	"errors"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/group/edwards25519"
)

var group = new(edwards25519.Curve)
var ErrPKMarshalling = errors.New("error unmarshalling public key")
var ErrPKInvalid = errors.New("invalid public key")
var ErrPKSmallOrder = errors.New("public key has small order")
var ErrPKNotCanonical = errors.New("public key is not canonical")

var ErrEdDSAWrongLength = errors.New("wrong length for decoding EdDSA private")
var ErrSchnorrInvalidScalar = errors.New("schnorr: s invalid scalar")
var ErrSignatureLength = errors.New("signature length invalid")
var ErrSignatureNotCanonical = errors.New("signature is not canonical")
var ErrSignatureRecNotEqual = errors.New("reconstructed S is not equal to signature")

var ErrPointRSmallOrder = errors.New("point R has small order")
var ErrPointRNotCanonical = errors.New("point R is not canonical")
var ErrPointRInvalid = errors.New("point R invalid")

// EdDSA is a structure holding the data necessary to make a series of
// EdDSA signatures.
type EdDSA struct {
	// Secret being already hashed + bit tweaked
	Secret kyber.Scalar
	// Public is the corresponding public key
	Public kyber.Point

	seed   []byte
	prefix []byte
}

// NewEdDSA will return a freshly generated key pair to use for generating
// EdDSA signatures.
func NewEdDSA(stream cipher.Stream) *EdDSA { _ = "STUB: not implemented"; return nil }

// MarshalBinary will return the representation used by the reference
// implementation of SUPERCOP ref10, which is "seed || Public".
func (e *EdDSA) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary transforms a slice of bytes into a EdDSA signature.
func (e *EdDSA) UnmarshalBinary(buff []byte) error { _ = "STUB: not implemented"; return nil }

// Sign will return a EdDSA signature of the message msg using Ed25519.
func (e *EdDSA) Sign(msg []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// deterministic random secret and its commit

// Compute challenge: H( R || Public || Msg)

// Compute response s = r + h * s

// return R || s

// VerifyWithChecks uses a public key buffer, a message and a signature.
// It will return nil if sig is a valid signature for msg created by
// key public, or an error otherwise. Compared to `Verify`, it performs
// additional checks around the canonicality and ensures the public key
// does not have a small order.
func VerifyWithChecks(pub, msg, sig []byte) error { _ = "STUB: not implemented"; return nil }

// reconstruct h = H(R || Public || Msg)

// reconstruct S == k*A + R

// Verify uses a public key, a message and a signature. It will return nil if
// sig is a valid signature for msg created by key public, or an error otherwise.
func Verify(public kyber.Point, msg, sig []byte) error { _ = "STUB: not implemented"; return nil }
