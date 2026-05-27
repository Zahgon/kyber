// Package bdn implements the Boneh-Drijvers-Neven signature scheme which is
// an implementation of the bls package which is robust against rogue public-key attacks. Those
// attacks could allow an attacker to forge a public-key and then make a verifiable
// signature for an aggregation of signatures. It fixes the situation by
// adding coefficients to the aggregate.
//
// See the papers:
// https://eprint.iacr.org/2018/483.pdf
// https://crypto.stanford.edu/~dabo/pubs/papers/BLSmultisig.html
package bdn

import (
	"crypto/cipher"

	"go.dedis.ch/kyber/v4"
	"go.dedis.ch/kyber/v4/pairing"
	"go.dedis.ch/kyber/v4/sign"
)

// For the choice of H, we're mostly worried about the second preimage attack. In
// other words, find m' where H(m) == H(m')
// We also use the entire roster so that the coefficient will vary for the same
// public key used in different roster
func hashPointToR(group kyber.Group, pubs []kyber.Point) ([]kyber.Scalar, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Scheme struct {
	blsScheme sign.Scheme
	sigGroup  kyber.Group
	keyGroup  kyber.Group
	pairing   func(signature, public, hashedPoint kyber.Point) bool
}

// NewSchemeOnG1 returns a sign.Scheme that uses G1 for its signature space and G2
// for its public keys
func NewSchemeOnG1(suite pairing.Suite) *Scheme { _ = "STUB: not implemented"; return nil }

// NewSchemeOnG2 returns a sign.Scheme that uses G2 for its signature space and
// G1 for its public key
func NewSchemeOnG2(suite pairing.Suite) *Scheme { _ = "STUB: not implemented"; return nil }

// NewKeyPair creates a new BLS signing key pair. The private key x is a scalar
// and the public key X is a point on the scheme's key group.
func (scheme *Scheme) NewKeyPair(random cipher.Stream) (kyber.Scalar, kyber.Point) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), *new(kyber.Point)
}

// Sign creates a BLS signature S = x * H(m) on a message m using the private
// key x. The signature S is a point on the scheme's signature group.
func (scheme *Scheme) Sign(x kyber.Scalar, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify checks the given BLS signature S on the message m using the public
// key X by verifying that the equality e(H(m), X) == e(H(m), x*B2) ==
// e(x*H(m), B2) == e(S, B2) holds where e is the pairing operation and B2 is
// the base point from the scheme's key group.
func (scheme *Scheme) Verify(x kyber.Point, msg, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// AggregateSignatures aggregates the signatures using a coefficient for each
// one of them where c = H(pk) and H: keyGroup -> R with R = {1, ..., 2^128}
func (scheme *Scheme) AggregateSignatures(sigs [][]byte, mask *Mask) (kyber.Point, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Point), nil
}

// this should never happen because of the loop boundary
// an error here is probably a bug in the mask implementation

// c+1 because R is in the range [1, 2^128] and not [0, 2^128-1]

// AggregatePublicKeys aggregates a set of public keys (similarly to
// AggregateSignatures for signatures) using the hash function
// H: keyGroup -> R with R = {1, ..., 2^128}.
func (scheme *Scheme) AggregatePublicKeys(mask *Mask) (kyber.Point, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Point), nil
}

// this should never happen because of the loop boundary
// an error here is probably a bug in the mask implementation

// v1 API Deprecated ----------------------------------

// NewKeyPair creates a new BLS signing key pair. The private key x is a scalar
// and the public key X is a point on curve G2.
//
// Deprecated: use the new scheme methods instead.
func NewKeyPair(suite pairing.Suite, random cipher.Stream) (kyber.Scalar, kyber.Point) {
	_ = "STUB: not implemented"
	return *new(kyber.Scalar), *new(kyber.Point)
}

// Sign creates a BLS signature S = x * H(m) on a message m using the private
// key x. The signature S is a point on curve G1.
//
// Deprecated: use the new scheme methods instead.
func Sign(suite pairing.Suite, x kyber.Scalar, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify checks the given BLS signature S on the message m using the public
// key X by verifying that the equality e(H(m), X) == e(H(m), x*B2) ==
// e(x*H(m), B2) == e(S, B2) holds where e is the pairing operation and B2 is
// the base point from curve G2.
//
// Deprecated: use the new scheme methods instead.
func Verify(suite pairing.Suite, x kyber.Point, msg, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// AggregateSignatures aggregates the signatures using a coefficient for each
// one of them where c = H(pk) and H: G2 -> R with R = {1, ..., 2^128}
//
// Deprecated: use the new scheme methods instead.
func AggregateSignatures(suite pairing.Suite, sigs [][]byte, mask *Mask) (kyber.Point, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Point), nil
}

// AggregatePublicKeys aggregates a set of public keys (similarly to
// AggregateSignatures for signatures) using the hash function
// H: G2 -> R with R = {1, ..., 2^128}.
//
// Deprecated: use the new scheme methods instead.
func AggregatePublicKeys(suite pairing.Suite, mask *Mask) (kyber.Point, error) {
	_ = "STUB: not implemented"
	return *new(kyber.Point), nil
}
