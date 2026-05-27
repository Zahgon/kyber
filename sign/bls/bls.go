// Package bls implements the Boneh-Lynn-Shacham (BLS) signature scheme which
// was introduced in the paper "Short Signatures from the Weil Pairing". BLS
// requires pairing-based cryptography.
//
// When using aggregated signatures, this version is vulnerable to rogue
// public-key attack.
// The `sign/bdn` package should be used to make sure a signature
// aggregate cannot be verified by a forged key. You can find the protocol
// in kyber/sign/bdn. Note that only the aggregation is broken against the
// attack and for that reason, the code performing aggregation was removed.
//
// See the paper: https://crypto.stanford.edu/~dabo/pubs/papers/BLSmultisig.html
package bls

import (
	"crypto/cipher"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
	"go.dedis.ch/kyber/v4/sign"
)

type scheme struct {
	sigGroup kyber.Group
	keyGroup kyber.Group
	pairing  func(signature, public, hashedPoint kyber.Point) bool
}

// NewSchemeOnG1 returns a sign.Scheme that uses G1 for its signature space and G2
// for its public keys
func NewSchemeOnG1(suite pairing.Suite) sign.Scheme {
	_ = "STUB: not implemented"
	return *new(sign.Scheme)
}

// NewSchemeOnG2 returns a sign.Scheme that uses G2 for its signature space and
// G1 for its public key
func NewSchemeOnG2(suite pairing.Suite) sign.Scheme {
	_ = "STUB: not implemented"
	return *new(sign.Scheme)
}

func (s *scheme) NewKeyPair(random cipher.Stream) (kyber.Scalar, kyber.Point) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), *new(kyber.Point)
}

func (s *scheme) Sign(private kyber.Scalar, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *scheme) Verify(X kyber.Point, msg, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}
