/*
Package schnorr implements the vanilla Schnorr signature scheme.
See https://en.wikipedia.org/wiki/Schnorr_signature.

The only difference regarding the vanilla reference is the computation of the
response. This implementation adds the random component with the challenge times
private key while the Wikipedia article subtracts them.

The resulting signature is compatible with EdDSA verification algorithm when
using the edwards25519 group, and by extension the CoSi verification algorithm.
*/
package schnorr

import (
	"crypto/cipher"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/sign"
)

// Suite represents the set of functionalities needed by the package schnorr.
type Suite interface {
	kyber.Group
	kyber.Random
}

type Scheme struct {
	s Suite
}

func NewScheme(s Suite) sign.Scheme { _ = "STUB: not implemented"; return *new(sign.Scheme) }

func (s *Scheme) NewKeyPair(random cipher.Stream) (kyber.Scalar, kyber.Point) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), *new(kyber.Point)
}

func (s *Scheme) Sign(private kyber.Scalar, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Scheme) Verify(public kyber.Point, msg, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Sign creates a Sign signature from a msg and a private key. This
// signature can be verified with VerifySchnorr. It's also a valid EdDSA
// signature when using the edwards25519 Group.
func Sign(s Suite, private kyber.Scalar, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// create random secret k and public point commitment R
		nil
}

// create hash(public || R || message)

// compute response s = k + x*h

// return R || s

// VerifyWithChecks uses a public key buffer, a message and a signature.
// It will return nil if sig is a valid signature for msg created by
// key public, or an error otherwise. Compared to `Verify`, it performs
// additional checks around the canonicality and ensures the public key
// does not have a small order when using `edwards25519` group.
func VerifyWithChecks(g kyber.Group, pub, msg, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// recompute hash(public || R || msg)

// compute S = g^s

// compute RAh = R + A^h

// Verify verifies a given Schnorr signature. It returns nil iff the
// given signature is valid.
func Verify(g kyber.Group, public kyber.Point, msg, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func hash(g kyber.Group, public, r kyber.Point, msg []byte) (kyber.Scalar, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), nil
}
