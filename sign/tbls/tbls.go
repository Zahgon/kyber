// Package tbls implements the (t,n)-threshold Boneh-Lynn-Shacham signature
// scheme. During setup a group of n participants runs a distributed key
// generation algorithm (see kyber/share/dkg) to compute a joint public signing
// key X and one secret key share xi for each of the n signers. To compute a
// signature S on a message m, at least t ouf of n signers have to provide
// partial (BLS) signatures Si on m using their individual key shares xi which
// can then be used to recover the full (regular) BLS signature S via Lagrange
// interpolation. The signature S can be verified with the initially
// established group key X. Signatures are points on curve G1 and public keys
// are points on curve G2.
package tbls

import (
	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
	"go.dedis.ch/kyber/v4/share"
	"go.dedis.ch/kyber/v4/sign"
)

// SigShare encodes a threshold BLS signature share Si = i || v where the 2-byte
// big-endian value i corresponds to the share's index and v represents the
// share's value. The signature share Si is a point on curve G1 or G2.
type SigShare []byte

// Index returns the index i of the TBLS share Si.
func (s SigShare) Index() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Value returns the value v of the TBLS share Si.
func (s *SigShare) Value() []byte { _ = "STUB: not implemented"; return nil }

type scheme struct {
	keyGroup kyber.Group
	sigGroup kyber.Group
	sign.Scheme
}

// NewThresholdSchemeOnG1 returns a treshold scheme that computes bls signatures
// on G1
func NewThresholdSchemeOnG1(suite pairing.Suite) sign.ThresholdScheme {
	_ = "STUB: not implemented"
	return *new(sign.ThresholdScheme)
}

// NewThresholdSchemeOnG2 returns a treshold scheme that computes bls signatures
// on G2
func NewThresholdSchemeOnG2(suite pairing.Suite) sign.ThresholdScheme {
	_ = "STUB: not implemented"
	return *new(sign.ThresholdScheme)
}

// Sign creates a threshold BLS signature Si = xi * H(m) on the given message m
// using the provided secret key share xi.
func (s *scheme) Sign(private *share.PriShare, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *scheme) IndexOf(signature []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// VerifyPartial checks the given threshold BLS signature Si on the message m using
// the public key share Xi that is associated to the secret key share xi. This
// public key share Xi can be computed by evaluating the public sharing
// polynonmial at the share's index i.
func (s *scheme) VerifyPartial(public *share.PubPoly, msg, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *scheme) VerifyRecovered(public kyber.Point, msg, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Recover reconstructs the full BLS signature S = x * H(m) from a threshold t
// of signature shares Si using Lagrange interpolation. The full signature S
// can be verified through the regular BLS verification routine using the
// shared public key X. The shared public key can be computed by evaluating the
// public sharing polynomial at index 0.
func (s *scheme) Recover(public *share.PubPoly, msg []byte, sigs [][]byte, t, n uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
